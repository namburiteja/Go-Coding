import { useCallback, useEffect, useState } from 'react'
import { getAdminById } from '../services/api/admin.api'
import type { AdminProfile } from '../types/admin'
import { getErrorMessage } from '../types/api'
import { useAuth } from '../context/useAuth'

export function useAdminProfile() {
  const { user } = useAuth()
  const [profile, setProfile] = useState<AdminProfile | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    if (!user) {
      setProfile(null)
      setError('Not authenticated')
      setLoading(false)
      return
    }

    setLoading(true)
    setError(null)
    try {
      const data = await getAdminById(user.userId)
      setProfile(data)
    } catch (err) {
      setProfile(null)
      setError(getErrorMessage(err, 'Failed to load admin profile'))
    } finally {
      setLoading(false)
    }
  }, [user])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      if (!user) {
        if (!cancelled) {
          setProfile(null)
          setError('Not authenticated')
          setLoading(false)
        }
        return
      }

      try {
        const data = await getAdminById(user.userId)
        if (!cancelled) {
          setProfile(data)
          setError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setProfile(null)
          setError(getErrorMessage(err, 'Failed to load admin profile'))
        }
      } finally {
        if (!cancelled) {
          setLoading(false)
        }
      }
    }

    void loadInitial()
    return () => {
      cancelled = true
    }
  }, [user])

  return { profile, loading, error, refresh, setProfile }
}

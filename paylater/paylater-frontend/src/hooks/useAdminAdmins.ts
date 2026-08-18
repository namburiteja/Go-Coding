import { useCallback, useEffect, useState } from 'react'
import { getAdmins } from '../services/api/admin.api'
import type { AdminProfile } from '../types/admin'
import { getErrorMessage } from '../types/api'

export function useAdminAdmins() {
  const [admins, setAdmins] = useState<AdminProfile[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await getAdmins()
      setAdmins(data)
    } catch (err) {
      setAdmins([])
      setError(getErrorMessage(err, 'Failed to load admins'))
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      try {
        const data = await getAdmins()
        if (!cancelled) {
          setAdmins(data)
          setError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setAdmins([])
          setError(getErrorMessage(err, 'Failed to load admins'))
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
  }, [])

  return { admins, loading, error, refresh, setAdmins }
}

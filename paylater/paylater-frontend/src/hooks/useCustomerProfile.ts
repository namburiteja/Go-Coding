import { useCallback, useEffect, useState } from 'react'
import { getCustomerProfile } from '../services/api/customer.api'
import type { CustomerProfile } from '../types/customer'
import { getErrorMessage } from '../types/api'

export function useCustomerProfile() {
  const [profile, setProfile] = useState<CustomerProfile | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await getCustomerProfile()
      setProfile(data)
    } catch (err) {
      setProfile(null)
      setError(getErrorMessage(err, 'Failed to load profile'))
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      try {
        const data = await getCustomerProfile()
        if (!cancelled) {
          setProfile(data)
          setError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setProfile(null)
          setError(getErrorMessage(err, 'Failed to load profile'))
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

  return { profile, loading, error, refresh, setProfile }
}

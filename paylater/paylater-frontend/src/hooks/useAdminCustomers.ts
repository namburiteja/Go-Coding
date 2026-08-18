import { useCallback, useEffect, useState } from 'react'
import { getCustomers } from '../services/api/admin.api'
import type { AdminCustomer } from '../types/admin'
import { getErrorMessage } from '../types/api'

export function useAdminCustomers() {
  const [customers, setCustomers] = useState<AdminCustomer[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await getCustomers()
      setCustomers(data)
    } catch (err) {
      setCustomers([])
      setError(getErrorMessage(err, 'Failed to load customers'))
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      try {
        const data = await getCustomers()
        if (!cancelled) {
          setCustomers(data)
          setError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setCustomers([])
          setError(getErrorMessage(err, 'Failed to load customers'))
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

  return { customers, loading, error, refresh, setCustomers }
}

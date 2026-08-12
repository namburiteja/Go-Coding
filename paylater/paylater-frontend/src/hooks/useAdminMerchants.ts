import { useCallback, useEffect, useState } from 'react'
import { getMerchants } from '../services/api/admin.api'
import type { AdminMerchant } from '../types/admin'
import { getErrorMessage } from '../types/api'

export function useAdminMerchants() {
  const [merchants, setMerchants] = useState<AdminMerchant[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await getMerchants()
      setMerchants(data)
    } catch (err) {
      setMerchants([])
      setError(getErrorMessage(err, 'Failed to load merchants'))
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      try {
        const data = await getMerchants()
        if (!cancelled) {
          setMerchants(data)
          setError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setMerchants([])
          setError(getErrorMessage(err, 'Failed to load merchants'))
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

  return { merchants, loading, error, refresh, setMerchants }
}

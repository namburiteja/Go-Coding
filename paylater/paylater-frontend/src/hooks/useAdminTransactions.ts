import { useCallback, useEffect, useState } from 'react'
import { getAllTransactions } from '../services/api/admin.api'
import type { AdminTransaction } from '../types/admin'
import { getErrorMessage } from '../types/api'

export function useAdminTransactions() {
  const [transactions, setTransactions] = useState<AdminTransaction[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await getAllTransactions()
      setTransactions(data)
    } catch (err) {
      setTransactions([])
      setError(getErrorMessage(err, 'Failed to load transactions'))
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      try {
        const data = await getAllTransactions()
        if (!cancelled) {
          setTransactions(data)
          setError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setTransactions([])
          setError(getErrorMessage(err, 'Failed to load transactions'))
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

  return { transactions, loading, error, refresh, setTransactions }
}

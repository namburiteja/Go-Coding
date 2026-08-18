import { useCallback, useEffect, useState } from 'react'
import { getMerchantProfile } from '../services/api/merchant.api'
import type { MerchantProfile } from '../types/merchant'
import { getErrorMessage } from '../types/api'

export function useMerchantProfile() {
  const [profile, setProfile] = useState<MerchantProfile | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await getMerchantProfile()
      setProfile(data)
    } catch (err) {
      setProfile(null)
      setError(getErrorMessage(err, 'Failed to load merchant profile'))
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      try {
        const data = await getMerchantProfile()
        if (!cancelled) {
          setProfile(data)
          setError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setProfile(null)
          setError(getErrorMessage(err, 'Failed to load merchant profile'))
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

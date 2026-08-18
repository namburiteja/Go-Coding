import { useState } from 'react'
import { getErrorMessage } from '../types/api'

export function useApiError() {
  const [error, setError] = useState<string | null>(null)

  function capture(err: unknown, fallback?: string) {
    setError(getErrorMessage(err, fallback))
  }

  function clear() {
    setError(null)
  }

  return { error, setError, capture, clear }
}

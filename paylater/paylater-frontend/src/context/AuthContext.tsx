import { useCallback, useMemo, useState, type ReactNode, } from 'react'
import type { AuthUser } from '../types/auth'
import { clearStoredToken, decodeJwt, getStoredToken, isTokenExpired, isValidRole, setStoredToken, } from '../utils/jwt'
import { AuthContext } from './auth-context'

function userFromToken(token: string): AuthUser | null {
  const claims = decodeJwt(token)
  if (!claims || !isValidRole(claims.role) || isTokenExpired(claims)) {
    return null
  }
  return {
    userId: claims.user_id,
    role: claims.role,
    token
  }
}

function readInitialUser(): AuthUser | null {
  const token = getStoredToken()
  if (!token) {
    return null
  }
  const user = userFromToken(token)
  if (!user) {
    clearStoredToken()
    return null
  }
  return user
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<AuthUser | null>(() => readInitialUser())

  const login = useCallback((token: string) => {
    const next = userFromToken(token)
    if (!next) {
      clearStoredToken()
      setUser(null)
      throw new Error('Invalid or expired authentication token')
    }
    setStoredToken(token)
    setUser(next)
  }, [])

  const logout = useCallback(() => {
    clearStoredToken()
    setUser(null)
  }, [])

  const value = useMemo(
    () => ({
      user,
      isAuthenticated: Boolean(user),
      role: user?.role ?? null,
      userId: user?.userId ?? null,
      login,
      logout,
    }),
    [user, login, logout],
  )

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

import { createContext } from 'react'
import type { AuthUser, Role } from '../types/auth'

export interface AuthContextValue {
  user: AuthUser | null
  isAuthenticated: boolean
  role: Role | null
  userId: number | null
  login: (token: string) => void
  logout: () => void
}

export const AuthContext = createContext<AuthContextValue | undefined>(undefined)

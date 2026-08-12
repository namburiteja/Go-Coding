import type { JwtClaims, Role } from '../types/auth'

const TOKEN_KEY = 'paylater_token'

export function getStoredToken(): string | null {
  return localStorage.getItem(TOKEN_KEY)
}

export function setStoredToken(token: string): void {
  localStorage.setItem(TOKEN_KEY, token)
}

export function clearStoredToken(): void {
  localStorage.removeItem(TOKEN_KEY)
}

function base64UrlDecode(input: string): string {
  const normalized = input.replace(/-/g, '+').replace(/_/g, '/')
  const padded = normalized.padEnd(normalized.length + ((4 - (normalized.length % 4)) % 4), '=')
  return atob(padded)
}

export function decodeJwt(token: string): JwtClaims | null {
  try {
    const parts = token.split('.')
    if (parts.length < 2) {
      return null
    }
    const payload = JSON.parse(base64UrlDecode(parts[1])) as JwtClaims
    if (!payload.user_id || !payload.role) {
      return null
    }
    return payload
  } catch {
    return null
  }
}

export function isTokenExpired(claims: JwtClaims): boolean {
  if (!claims.exp) {
    return false
  }
  return claims.exp * 1000 <= Date.now()
}

export function isValidRole(role: string): role is Role {
  return role === 'ADMIN' || role === 'MERCHANT' || role === 'CUSTOMER'
}

export function dashboardPathForRole(role: Role): string {
  switch (role) {
    case 'ADMIN':
      return '/admin/dashboard'
    case 'MERCHANT':
      return '/merchant/dashboard'
    case 'CUSTOMER':
      return '/customer/dashboard'
  }
}

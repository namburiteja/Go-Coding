export type Role = 'ADMIN' | 'MERCHANT' | 'CUSTOMER'

export interface JwtClaims {
  user_id: number
  role: Role
  exp?: number
  iat?: number
}

export interface AuthUser {
  userId: number
  role: Role
  token: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface LoginResponse {
  token: string
}

export interface CustomerRegisterRequest {
  name: string
  email: string
  password: string
}

export interface MerchantRegisterRequest {
  name: string
  email: string
  phone: string
  password: string
}

export interface MessageResponse {
  message: string
}

export interface ApiErrorBody {
  error: string
}

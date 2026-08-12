import axios, { type AxiosError } from 'axios'
import { ApiError } from '../../types/api'
import type { ApiErrorBody } from '../../types/auth'
import { clearStoredToken, getStoredToken } from '../../utils/jwt'

const baseURL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:9090'

export const apiClient = axios.create({
  baseURL,
  headers: {
    'Content-Type': 'application/json',
  },
})

apiClient.interceptors.request.use((config) => {
  const token = getStoredToken()
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

apiClient.interceptors.response.use(
  (response) => response,
  (error: AxiosError<ApiErrorBody>) => {
    const status = error.response?.status ?? 0
    const message =
      error.response?.data?.error ||
      error.message ||
      'Request failed'

    if (status === 401) {
      clearStoredToken()
      if (typeof window !== 'undefined' && !window.location.pathname.includes('/login')) {
        const path = window.location.pathname
        let loginPath = '/customer/login'
        if (path.startsWith('/admin')) {
          loginPath = '/admin/login'
        } else if (path.startsWith('/merchant')) {
          loginPath = '/merchant/login'
        }
        window.location.assign(loginPath)
      }
    }

    return Promise.reject(new ApiError(message, status))
  },
)

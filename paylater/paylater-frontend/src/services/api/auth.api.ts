import type {
  CustomerRegisterRequest,
  LoginRequest,
  LoginResponse,
  MerchantRegisterRequest,
  MessageResponse,
} from '../../types/auth'
import { apiClient } from './client'

export async function loginCustomer(payload: LoginRequest): Promise<LoginResponse> {
  const { data } = await apiClient.post<LoginResponse>('/customers/login', payload)
  return data
}

export async function registerCustomer(payload: CustomerRegisterRequest): Promise<MessageResponse> {
  const { data } = await apiClient.post<MessageResponse>('/customers/register', payload)
  return data
}

export async function loginMerchant(payload: LoginRequest): Promise<LoginResponse> {
  const { data } = await apiClient.post<LoginResponse>('/merchants/login', payload)
  return data
}

export async function registerMerchant(payload: MerchantRegisterRequest): Promise<MessageResponse> {
  const { data } = await apiClient.post<MessageResponse>('/merchants/register', payload)
  return data
}

export async function loginAdmin(payload: LoginRequest): Promise<LoginResponse> {
  const { data } = await apiClient.post<LoginResponse>('/admins/login', payload)
  return data
}

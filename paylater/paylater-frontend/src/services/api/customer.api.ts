import type { MessageResponse } from '../../types/auth'
import type {
  CustomerProfile,
  CustomerTransaction,
  MerchantOption,
  PaybackRequest,
  PurchaseRequest,
  UpdateCustomerProfileRequest,
} from '../../types/customer'
import { apiClient } from './client'

export async function getCustomerProfile(): Promise<CustomerProfile> {
  const { data } = await apiClient.get<CustomerProfile>('/customers/me')
  return data
}

export async function updateCustomerProfile(
  payload: UpdateCustomerProfileRequest,
): Promise<MessageResponse> {
  const { data } = await apiClient.put<MessageResponse>('/customers/me', payload)
  return data
}

export async function getCustomerTransactions(): Promise<CustomerTransaction[]> {
  const { data } = await apiClient.get<CustomerTransaction[]>('/customers/me/transactions')
  return data
}

export async function purchase(payload: PurchaseRequest): Promise<MessageResponse> {
  const { data } = await apiClient.post<MessageResponse>('/customers/purchase', payload)
  return data
}

export async function payback(payload: PaybackRequest): Promise<MessageResponse> {
  const { data } = await apiClient.post<MessageResponse>('/customers/payback', payload)
  return data
}

export async function getMerchantOptions(): Promise<MerchantOption[]> {
  const { data } = await apiClient.get<MerchantOption[]>('/merchants/options')
  return data
}

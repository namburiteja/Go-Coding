import type { MessageResponse } from '../../types/auth'
import type {
  MerchantProfile,
  MerchantTransaction,
  UpdateMerchantProfileRequest,
} from '../../types/merchant'
import { apiClient } from './client'

export async function getMerchantProfile(): Promise<MerchantProfile> {
  const { data } = await apiClient.get<MerchantProfile>('/merchants/me')
  return data
}

export async function updateMerchantProfile(
  payload: UpdateMerchantProfileRequest,
): Promise<MessageResponse> {
  const { data } = await apiClient.put<MessageResponse>('/merchants/me', payload)
  return data
}

export async function getMerchantTransactions(): Promise<MerchantTransaction[]> {
  const { data } = await apiClient.get<MerchantTransaction[]>('/merchants/me/transactions')
  return data
}

import type { MessageResponse } from '../../types/auth'
import type {
  AdminCustomer,
  AdminMerchant,
  AdminProfile,
  AdminRegisterRequest,
  AdminTransaction,
  MerchantFeeReport,
  UpdateAdminCustomerRequest,
  UpdateAdminMerchantRequest,
  UpdateAdminRequest,
  UpdateCommissionRequest,
} from '../../types/admin'
import { apiClient } from './client'

// —— Admins ——

export async function getAdmins(): Promise<AdminProfile[]> {
  const { data } = await apiClient.get<AdminProfile[]>('/admins')
  return data
}

export async function getAdminById(id: number): Promise<AdminProfile> {
  const { data } = await apiClient.get<AdminProfile>(`/admins/${id}`)
  return data
}

export async function registerAdmin(payload: AdminRegisterRequest): Promise<MessageResponse> {
  const { data } = await apiClient.post<MessageResponse>('/admins/register', payload)
  return data
}

export async function updateAdmin(
  id: number,
  payload: UpdateAdminRequest,
): Promise<MessageResponse> {
  const { data } = await apiClient.put<MessageResponse>(`/admins/${id}`, payload)
  return data
}

export async function deleteAdmin(id: number): Promise<MessageResponse> {
  const { data } = await apiClient.delete<MessageResponse>(`/admins/${id}`)
  return data
}

// —— Customers (admin) ——

export async function getCustomers(): Promise<AdminCustomer[]> {
  const { data } = await apiClient.get<AdminCustomer[]>('/customers')
  return data
}

export async function getCustomerById(id: number): Promise<AdminCustomer> {
  const { data } = await apiClient.get<AdminCustomer>(`/customers/${id}`)
  return data
}

export async function updateCustomer(
  id: number,
  payload: UpdateAdminCustomerRequest,
): Promise<MessageResponse> {
  const { data } = await apiClient.put<MessageResponse>(`/customers/${id}`, payload)
  return data
}

export async function deleteCustomer(id: number): Promise<MessageResponse> {
  const { data } = await apiClient.delete<MessageResponse>(`/customers/${id}`)
  return data
}

// —— Merchants (admin) ——

export async function getMerchants(): Promise<AdminMerchant[]> {
  const { data } = await apiClient.get<AdminMerchant[]>('/merchants')
  return data
}

export async function getMerchantById(id: number): Promise<AdminMerchant> {
  const { data } = await apiClient.get<AdminMerchant>(`/merchants/${id}`)
  return data
}

export async function updateMerchant(
  id: number,
  payload: UpdateAdminMerchantRequest,
): Promise<MessageResponse> {
  const { data } = await apiClient.put<MessageResponse>(`/merchants/${id}`, payload)
  return data
}

export async function updateMerchantCommission(
  id: number,
  payload: UpdateCommissionRequest,
): Promise<MessageResponse> {
  const { data } = await apiClient.put<MessageResponse>(
    `/merchants/${id}/commission`,
    payload,
  )
  return data
}

export async function deleteMerchant(id: number): Promise<MessageResponse> {
  const { data } = await apiClient.delete<MessageResponse>(`/merchants/${id}`)
  return data
}

// —— Transactions (admin) ——

export async function getAllTransactions(): Promise<AdminTransaction[]> {
  const { data } = await apiClient.get<AdminTransaction[]>('/transactions')
  return data
}

// —— Reports (admin) ——

export async function getUsersAtCreditLimit(): Promise<AdminCustomer[]> {
  const { data } = await apiClient.get<AdminCustomer[]>('/reports/credit-limit')
  return data
}

export async function getCustomersWithDue(): Promise<AdminCustomer[]> {
  const { data } = await apiClient.get<AdminCustomer[]>('/reports/customers-due')
  return data
}

export async function getCustomerDueByName(name: string): Promise<AdminCustomer> {
  const { data } = await apiClient.get<AdminCustomer>(
    `/reports/customer-due/${encodeURIComponent(name)}`,
  )
  return data
}

export async function getMerchantFees(): Promise<MerchantFeeReport[]> {
  const { data } = await apiClient.get<MerchantFeeReport[]>('/reports/merchant-fees')
  return data
}

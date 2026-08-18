import type { CustomerProfile } from './customer'
import type { MerchantProfile } from './merchant'

/** Matches GET /admins and GET /admins/:id — password/hash never included. */
export interface AdminProfile {
  id: number
  name: string
  email: string
  role: string
  created_at?: string | null
}

/** POST /admins/register (Admin JWT required). */
export interface AdminRegisterRequest {
  name: string
  email: string
  password: string
}

export interface UpdateAdminRequest {
  name: string
  email: string
}

/** Admin customer list/detail uses the same public customer shape. */
export type AdminCustomer = CustomerProfile

export interface UpdateAdminCustomerRequest {
  name: string
  email: string
}

/** Admin merchant list/detail uses the same public merchant shape. */
export type AdminMerchant = MerchantProfile

export interface UpdateAdminMerchantRequest {
  name: string
  email: string
}

/** PUT /merchants/:id/commission */
export interface UpdateCommissionRequest {
  commission_percentage: string
}

/** Matches GET /transactions (ledger public shape). */
export interface AdminTransaction {
  id: number
  customer_id: number
  merchant_id: number | null
  transaction_type: string
  amount: string
  commission_percentage: string | null
  commission_amount: string | null
  transaction_date: string | null
}

/** Matches GET /reports/merchant-fees */
export interface MerchantFeeReport {
  id: number
  name: string
  total_fee_collected: string
}

export interface AdminDashboardStats {
  customerCount: number
  merchantCount: number
  adminCount: number
  transactionCount: number
  purchaseCount: number
  paybackCount: number
  totalPurchaseVolume: number
  totalFees: number
  blockedCustomers: number
  customersWithDue: number
}

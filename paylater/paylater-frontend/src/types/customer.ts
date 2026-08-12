export type CustomerStatus = 'ACTIVE' | 'BLOCKED'

/** Matches GET /customers/me (snake_case from backend). */
export interface CustomerProfile {
  id: number
  name: string
  email: string
  credit_limit: string
  total_due: string | null
  payment_due_date: string
  status: CustomerStatus | string | null
  created_at?: string | null
}

export interface UpdateCustomerProfileRequest {
  name: string
  email: string
}

export type TransactionType = 'PURCHASE' | 'PAYBACK'

/** Matches GET /customers/me/transactions. */
export interface CustomerTransaction {
  id: number
  customer_id: number
  merchant_id: number | null
  transaction_type: TransactionType | string
  amount: string
  commission_percentage: string | null
  commission_amount: string | null
  transaction_date: string | null
}

export interface PurchaseRequest {
  merchantId: number
  amount: string
}

export interface PaybackRequest {
  amount: string
}

/** Matches GET /merchants/options. */
export interface MerchantOption {
  id: number
  name: string
}

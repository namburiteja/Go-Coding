/** Matches GET /merchants/me (snake_case from backend). */
export interface MerchantProfile {
  id: number
  name: string
  email: string
  phone: string
  commission_percentage: string | null
  created_at?: string | null
}

export interface UpdateMerchantProfileRequest {
  name: string
  email: string
}

/** Matches GET /merchants/me/transactions (same ledger shape as customer txs). */
export interface MerchantTransaction {
  id: number
  customer_id: number
  merchant_id: number | null
  transaction_type: string
  amount: string
  commission_percentage: string | null
  commission_amount: string | null
  transaction_date: string | null
}

export interface MerchantSettlementSummary {
  saleCount: number
  totalSales: number
  totalCommission: number
  netSettlement: number
  uniqueCustomers: number
}

export interface DerivedMerchantCustomer {
  customerId: number
  purchaseCount: number
  totalSpent: number
  totalCommission: number
  lastPurchaseAt: string | null
}

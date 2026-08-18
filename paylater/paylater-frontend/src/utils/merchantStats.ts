import type { DerivedMerchantCustomer, MerchantSettlementSummary, MerchantTransaction, } from '../types/merchant'
import { parseAmount } from './credit'

/** Settlement figures derived from merchant purchase transactions (no separate settlement API). */
export function summarizeMerchantSettlement(
  transactions: MerchantTransaction[],
): MerchantSettlementSummary {
  const purchases = transactions.filter((tx) => tx.transaction_type === 'PURCHASE')
  let totalSales = 0
  let totalCommission = 0
  const customers = new Set<number>()

  for (const tx of purchases) {
    totalSales += parseAmount(tx.amount)
    totalCommission += parseAmount(tx.commission_amount)
    customers.add(tx.customer_id)
  }

  return {
    saleCount: purchases.length,
    totalSales,
    totalCommission,
    netSettlement: totalSales - totalCommission,
    uniqueCustomers: customers.size,
  }
}

/** Unique customers derived from this merchant's purchase history. */
export function deriveMerchantCustomers(
  transactions: MerchantTransaction[],
): DerivedMerchantCustomer[] {
  const map = new Map<number, DerivedMerchantCustomer>()

  for (const tx of transactions) {
    if (tx.transaction_type !== 'PURCHASE') continue

    const existing = map.get(tx.customer_id)
    const amount = parseAmount(tx.amount)
    const commission = parseAmount(tx.commission_amount)

    if (!existing) {
      map.set(tx.customer_id, {
        customerId: tx.customer_id,
        purchaseCount: 1,
        totalSpent: amount,
        totalCommission: commission,
        lastPurchaseAt: tx.transaction_date,
      })
      continue
    }

    existing.purchaseCount += 1
    existing.totalSpent += amount
    existing.totalCommission += commission
    if (
      tx.transaction_date &&
      (!existing.lastPurchaseAt || tx.transaction_date > existing.lastPurchaseAt)
    ) {
      existing.lastPurchaseAt = tx.transaction_date
    }
  }

  return Array.from(map.values()).sort((a, b) => b.totalSpent - a.totalSpent)
}

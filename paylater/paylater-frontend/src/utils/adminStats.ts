import type { AdminCustomer, AdminDashboardStats, AdminTransaction } from '../types/admin'
import { parseAmount } from './credit'

export function summarizeAdminDashboard(
  customers: AdminCustomer[],
  merchantCount: number,
  adminCount: number,
  transactions: AdminTransaction[],
): AdminDashboardStats {
  let purchaseCount = 0
  let paybackCount = 0
  let totalPurchaseVolume = 0
  let totalFees = 0

  for (const tx of transactions) {
    if (tx.transaction_type === 'PURCHASE') {
      purchaseCount += 1
      totalPurchaseVolume += parseAmount(tx.amount)
      totalFees += parseAmount(tx.commission_amount)
    } else if (tx.transaction_type === 'PAYBACK') {
      paybackCount += 1
    }
  }

  const blockedCustomers = customers.filter(
    (c) => (c.status ?? '').toUpperCase() === 'BLOCKED',
  ).length
  const customersWithDue = customers.filter((c) => parseAmount(c.total_due) > 0).length

  return {
    customerCount: customers.length,
    merchantCount,
    adminCount,
    transactionCount: transactions.length,
    purchaseCount,
    paybackCount,
    totalPurchaseVolume,
    totalFees,
    blockedCustomers,
    customersWithDue,
  }
}

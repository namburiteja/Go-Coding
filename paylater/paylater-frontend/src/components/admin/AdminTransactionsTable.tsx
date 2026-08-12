import type { AdminTransaction } from '../../types/admin'
import { formatCurrency, formatDateTime } from '../../utils/credit'
import { EmptyState } from '../ui/EmptyState'

interface AdminTransactionsTableProps {
  transactions: AdminTransaction[]
  customerNames?: Record<number, string>
  merchantNames?: Record<number, string>
  emptyTitle?: string
  emptyDescription?: string
}

export function AdminTransactionsTable({
  transactions,
  customerNames = {},
  merchantNames = {},
  emptyTitle = 'No transactions',
  emptyDescription = 'Platform purchases and paybacks will show up here.',
}: AdminTransactionsTableProps) {
  if (transactions.length === 0) {
    return <EmptyState title={emptyTitle} description={emptyDescription} />
  }

  return (
    <div className="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-slate-200 text-left text-sm">
          <thead className="bg-slate-50 text-xs uppercase tracking-wide text-slate-500">
            <tr>
              <th className="px-4 py-3 font-medium">ID</th>
              <th className="px-4 py-3 font-medium">Type</th>
              <th className="px-4 py-3 font-medium">Customer</th>
              <th className="px-4 py-3 font-medium">Merchant</th>
              <th className="px-4 py-3 font-medium">Amount</th>
              <th className="px-4 py-3 font-medium">Fee %</th>
              <th className="px-4 py-3 font-medium">Fee</th>
              <th className="px-4 py-3 font-medium">Date</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {transactions.map((tx) => {
              const isPurchase = tx.transaction_type === 'PURCHASE'
              const customerLabel =
                customerNames[tx.customer_id] ?? `Customer #${tx.customer_id}`
              const merchantLabel =
                tx.merchant_id == null
                  ? '—'
                  : merchantNames[tx.merchant_id] ?? `Merchant #${tx.merchant_id}`

              return (
                <tr key={tx.id} className="hover:bg-slate-50/80">
                  <td className="px-4 py-3 text-slate-700">#{tx.id}</td>
                  <td className="px-4 py-3">
                    <span
                      className={`inline-flex rounded-full px-2 py-0.5 text-xs font-semibold ${
                        isPurchase
                          ? 'bg-sky-100 text-sky-800'
                          : 'bg-violet-100 text-violet-800'
                      }`}
                    >
                      {tx.transaction_type}
                    </span>
                  </td>
                  <td className="px-4 py-3 text-slate-700">{customerLabel}</td>
                  <td className="px-4 py-3 text-slate-700">{merchantLabel}</td>
                  <td className="px-4 py-3 font-medium text-slate-900">
                    {formatCurrency(tx.amount)}
                  </td>
                  <td className="px-4 py-3 text-slate-700">
                    {tx.commission_percentage != null
                      ? `${tx.commission_percentage}%`
                      : '—'}
                  </td>
                  <td className="px-4 py-3 text-slate-700">
                    {tx.commission_amount != null
                      ? formatCurrency(tx.commission_amount)
                      : '—'}
                  </td>
                  <td className="px-4 py-3 text-slate-600">
                    {formatDateTime(tx.transaction_date)}
                  </td>
                </tr>
              )
            })}
          </tbody>
        </table>
      </div>
    </div>
  )
}

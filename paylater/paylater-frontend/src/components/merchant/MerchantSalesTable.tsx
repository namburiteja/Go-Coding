import type { MerchantTransaction } from '../../types/merchant'
import { formatCurrency, formatDateTime } from '../../utils/credit'
import { EmptyState } from '../ui/EmptyState'

interface MerchantSalesTableProps {
  transactions: MerchantTransaction[]
  emptyTitle?: string
  emptyDescription?: string
}

export function MerchantSalesTable({
  transactions,
  emptyTitle = 'No sales yet',
  emptyDescription = 'PayLater purchases at your store will show up here.',
}: MerchantSalesTableProps) {
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
              <th className="px-4 py-3 font-medium">Customer</th>
              <th className="px-4 py-3 font-medium">Amount</th>
              <th className="px-4 py-3 font-medium">Commission %</th>
              <th className="px-4 py-3 font-medium">Fee</th>
              <th className="px-4 py-3 font-medium">Date</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {transactions.map((tx) => (
              <tr key={tx.id} className="hover:bg-slate-50/80">
                <td className="px-4 py-3 text-slate-700">#{tx.id}</td>
                <td className="px-4 py-3 text-slate-700">#{tx.customer_id}</td>
                <td className="px-4 py-3 font-medium text-slate-900">
                  {formatCurrency(tx.amount)}
                </td>
                <td className="px-4 py-3 text-slate-700">
                  {tx.commission_percentage != null ? `${tx.commission_percentage}%` : '—'}
                </td>
                <td className="px-4 py-3 text-slate-700">
                  {formatCurrency(tx.commission_amount)}
                </td>
                <td className="px-4 py-3 text-slate-600">
                  {formatDateTime(tx.transaction_date)}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

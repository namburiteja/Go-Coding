import type { DerivedMerchantCustomer } from '../../types/merchant'
import { formatCurrency, formatDateTime } from '../../utils/credit'
import { EmptyState } from '../ui/EmptyState'

interface MerchantCustomersTableProps {
  customers: DerivedMerchantCustomer[]
}

export function MerchantCustomersTable({ customers }: MerchantCustomersTableProps) {
  if (customers.length === 0) {
    return (
      <EmptyState
        title="No customers yet"
        description="Customers appear here after they purchase with PayLater at your store."
      />
    )
  }

  return (
    <div className="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-slate-200 text-left text-sm">
          <thead className="bg-slate-50 text-xs uppercase tracking-wide text-slate-500">
            <tr>
              <th className="px-4 py-3 font-medium">Customer ID</th>
              <th className="px-4 py-3 font-medium">Purchases</th>
              <th className="px-4 py-3 font-medium">Total spent</th>
              <th className="px-4 py-3 font-medium">Commission</th>
              <th className="px-4 py-3 font-medium">Last purchase</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {customers.map((row) => (
              <tr key={row.customerId} className="hover:bg-slate-50/80">
                <td className="px-4 py-3 font-medium text-slate-900">#{row.customerId}</td>
                <td className="px-4 py-3 text-slate-700">{row.purchaseCount}</td>
                <td className="px-4 py-3 text-slate-900">{formatCurrency(row.totalSpent)}</td>
                <td className="px-4 py-3 text-slate-700">
                  {formatCurrency(row.totalCommission)}
                </td>
                <td className="px-4 py-3 text-slate-600">
                  {formatDateTime(row.lastPurchaseAt)}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

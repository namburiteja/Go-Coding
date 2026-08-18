import type { AdminCustomer } from '../../types/admin'
import { formatCurrency, formatDate } from '../../utils/credit'
import { EmptyState } from '../ui/EmptyState'
import { StatusBadge } from '../ui/StatusBadge'
import { Button } from '../ui/Button'

interface CustomersTableProps {
  customers: AdminCustomer[]
  onEdit: (customer: AdminCustomer) => void
  onDelete: (customer: AdminCustomer) => void
  deletingId?: number | null
}

export function CustomersTable({
  customers,
  onEdit,
  onDelete,
  deletingId = null,
}: CustomersTableProps) {
  if (customers.length === 0) {
    return (
      <EmptyState
        title="No customers"
        description="Registered customers will appear here."
      />
    )
  }

  return (
    <div className="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-slate-200 text-left text-sm">
          <thead className="bg-slate-50 text-xs uppercase tracking-wide text-slate-500">
            <tr>
              <th className="px-4 py-3 font-medium">ID</th>
              <th className="px-4 py-3 font-medium">Name</th>
              <th className="px-4 py-3 font-medium">Email</th>
              <th className="px-4 py-3 font-medium">Credit limit</th>
              <th className="px-4 py-3 font-medium">Due</th>
              <th className="px-4 py-3 font-medium">Status</th>
              <th className="px-4 py-3 font-medium">Due date</th>
              <th className="px-4 py-3 font-medium">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {customers.map((customer) => (
              <tr key={customer.id} className="hover:bg-slate-50/80">
                <td className="px-4 py-3 text-slate-700">#{customer.id}</td>
                <td className="px-4 py-3 font-medium text-slate-900">{customer.name}</td>
                <td className="px-4 py-3 text-slate-700">{customer.email}</td>
                <td className="px-4 py-3 text-slate-700">
                  {formatCurrency(customer.credit_limit)}
                </td>
                <td className="px-4 py-3 text-slate-700">
                  {formatCurrency(customer.total_due)}
                </td>
                <td className="px-4 py-3">
                  {customer.status ? <StatusBadge status={customer.status} /> : '—'}
                </td>
                <td className="px-4 py-3 text-slate-600">
                  {formatDate(customer.payment_due_date)}
                </td>
                <td className="px-4 py-3">
                  <div className="flex flex-wrap gap-2">
                    <Button variant="secondary" size="sm" onClick={() => onEdit(customer)}>
                      Edit
                    </Button>
                    <Button
                      variant="ghost"
                      size="sm"
                      className="text-red-700 hover:bg-red-50"
                      loading={deletingId === customer.id}
                      onClick={() => onDelete(customer)}
                    >
                      Delete
                    </Button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

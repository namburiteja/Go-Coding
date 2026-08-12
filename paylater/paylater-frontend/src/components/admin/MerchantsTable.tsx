import type { AdminMerchant } from '../../types/admin'
import { formatDate } from '../../utils/credit'
import { EmptyState } from '../ui/EmptyState'
import { Button } from '../ui/Button'

interface MerchantsTableProps {
  merchants: AdminMerchant[]
  onEdit: (merchant: AdminMerchant) => void
  onDelete: (merchant: AdminMerchant) => void
  deletingId?: number | null
}

export function MerchantsTable({
  merchants,
  onEdit,
  onDelete,
  deletingId = null,
}: MerchantsTableProps) {
  if (merchants.length === 0) {
    return (
      <EmptyState
        title="No merchants"
        description="Registered merchants will appear here."
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
              <th className="px-4 py-3 font-medium">Phone</th>
              <th className="px-4 py-3 font-medium">Commission</th>
              <th className="px-4 py-3 font-medium">Created</th>
              <th className="px-4 py-3 font-medium">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {merchants.map((merchant) => (
              <tr key={merchant.id} className="hover:bg-slate-50/80">
                <td className="px-4 py-3 text-slate-700">#{merchant.id}</td>
                <td className="px-4 py-3 font-medium text-slate-900">{merchant.name}</td>
                <td className="px-4 py-3 text-slate-700">{merchant.email}</td>
                <td className="px-4 py-3 text-slate-700">{merchant.phone}</td>
                <td className="px-4 py-3 text-slate-700">
                  {merchant.commission_percentage != null
                    ? `${merchant.commission_percentage}%`
                    : '—'}
                </td>
                <td className="px-4 py-3 text-slate-600">
                  {formatDate(merchant.created_at)}
                </td>
                <td className="px-4 py-3">
                  <div className="flex flex-wrap gap-2">
                    <Button variant="secondary" size="sm" onClick={() => onEdit(merchant)}>
                      Edit
                    </Button>
                    <Button
                      variant="ghost"
                      size="sm"
                      className="text-red-700 hover:bg-red-50"
                      loading={deletingId === merchant.id}
                      onClick={() => onDelete(merchant)}
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

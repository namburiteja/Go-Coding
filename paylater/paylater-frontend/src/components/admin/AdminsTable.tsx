import type { AdminProfile } from '../../types/admin'
import { formatDate } from '../../utils/credit'
import { EmptyState } from '../ui/EmptyState'
import { Button } from '../ui/Button'

interface AdminsTableProps {
  admins: AdminProfile[]
  currentAdminId?: number
  onEdit: (admin: AdminProfile) => void
  onDelete: (admin: AdminProfile) => void
  deletingId?: number | null
}

export function AdminsTable({
  admins,
  currentAdminId,
  onEdit,
  onDelete,
  deletingId = null,
}: AdminsTableProps) {
  if (admins.length === 0) {
    return (
      <EmptyState title="No admins" description="Administrators will appear here." />
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
              <th className="px-4 py-3 font-medium">Role</th>
              <th className="px-4 py-3 font-medium">Created</th>
              <th className="px-4 py-3 font-medium">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {admins.map((admin) => {
              const isSelf = currentAdminId === admin.id
              return (
                <tr key={admin.id} className="hover:bg-slate-50/80">
                  <td className="px-4 py-3 text-slate-700">#{admin.id}</td>
                  <td className="px-4 py-3 font-medium text-slate-900">
                    {admin.name}
                    {isSelf ? (
                      <span className="ml-2 text-xs font-normal text-slate-500">(you)</span>
                    ) : null}
                  </td>
                  <td className="px-4 py-3 text-slate-700">{admin.email}</td>
                  <td className="px-4 py-3 text-slate-700">{admin.role}</td>
                  <td className="px-4 py-3 text-slate-600">{formatDate(admin.created_at)}</td>
                  <td className="px-4 py-3">
                    <div className="flex flex-wrap gap-2">
                      <Button variant="secondary" size="sm" onClick={() => onEdit(admin)}>
                        Edit
                      </Button>
                      <Button
                        variant="ghost"
                        size="sm"
                        className="text-red-700 hover:bg-red-50"
                        loading={deletingId === admin.id}
                        disabled={isSelf}
                        title={isSelf ? 'You cannot delete your own account here' : undefined}
                        onClick={() => onDelete(admin)}
                      >
                        Delete
                      </Button>
                    </div>
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

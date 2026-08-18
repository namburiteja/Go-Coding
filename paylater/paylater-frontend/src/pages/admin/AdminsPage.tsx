import { useState } from 'react'
import { AdminProfileForm } from '../../components/admin/AdminProfileForm'
import { AdminsTable } from '../../components/admin/AdminsTable'
import { CreateAdminForm } from '../../components/admin/CreateAdminForm'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useAuth } from '../../context/useAuth'
import { useAdminAdmins } from '../../hooks/useAdminAdmins'
import { deleteAdmin } from '../../services/api/admin.api'
import type { AdminProfile } from '../../types/admin'
import { getErrorMessage } from '../../types/api'

export function AdminAdminsPage() {
  const { user } = useAuth()
  const { admins, loading, error, refresh } = useAdminAdmins()
  const [selected, setSelected] = useState<AdminProfile | null>(null)
  const [actionError, setActionError] = useState<string | null>(null)
  const [actionSuccess, setActionSuccess] = useState<string | null>(null)
  const [deletingId, setDeletingId] = useState<number | null>(null)

  async function handleDelete(admin: AdminProfile) {
    if (user?.userId === admin.id) {
      setActionError('You cannot delete your own admin account while signed in.')
      return
    }

    const confirmed = window.confirm(
      `Delete admin "${admin.name}" (#${admin.id})? This cannot be undone.`,
    )
    if (!confirmed) return

    setActionError(null)
    setActionSuccess(null)
    setDeletingId(admin.id)
    try {
      const result = await deleteAdmin(admin.id)
      setActionSuccess(result.message || 'Admin deleted successfully')
      if (selected?.id === admin.id) {
        setSelected(null)
      }
      await refresh()
    } catch (err) {
      setActionError(getErrorMessage(err, 'Failed to delete admin'))
    } finally {
      setDeletingId(null)
    }
  }

  return (
    <div>
      <PageHeader
        title="Admins"
        description="Create and manage administrator accounts. Registration stays behind Admin JWT."
        actions={
          <Button variant="secondary" size="sm" onClick={() => void refresh()} disabled={loading}>
            Refresh
          </Button>
        }
      />

      {error ? <Alert variant="error">{error}</Alert> : null}
      {actionError ? <Alert variant="error">{actionError}</Alert> : null}
      {actionSuccess ? <Alert variant="success">{actionSuccess}</Alert> : null}

      <div className="mb-6">
        <CreateAdminForm
          onSuccess={() => {
            void refresh()
          }}
        />
      </div>

      {selected ? (
        <div className="mb-6">
          <AdminProfileForm
            key={`${selected.id}-${selected.name}-${selected.email}`}
            profile={selected}
            title="Edit admin"
            description="Update name and email for this administrator."
            onCancel={() => setSelected(null)}
            onSuccess={() => {
              void refresh()
            }}
          />
        </div>
      ) : null}

      {loading ? (
        <Spinner label="Loading admins…" />
      ) : (
        <AdminsTable
          admins={admins}
          currentAdminId={user?.userId}
          onEdit={setSelected}
          onDelete={(admin) => void handleDelete(admin)}
          deletingId={deletingId}
        />
      )}
    </div>
  )
}

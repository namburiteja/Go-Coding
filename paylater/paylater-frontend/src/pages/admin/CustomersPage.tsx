import { useState } from 'react'
import { CustomerEditForm } from '../../components/admin/CustomerEditForm'
import { CustomersTable } from '../../components/admin/CustomersTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useAdminCustomers } from '../../hooks/useAdminCustomers'
import { deleteCustomer } from '../../services/api/admin.api'
import type { AdminCustomer } from '../../types/admin'
import { getErrorMessage } from '../../types/api'

export function AdminCustomersPage() {
  const { customers, loading, error, refresh } = useAdminCustomers()
  const [selected, setSelected] = useState<AdminCustomer | null>(null)
  const [actionError, setActionError] = useState<string | null>(null)
  const [actionSuccess, setActionSuccess] = useState<string | null>(null)
  const [deletingId, setDeletingId] = useState<number | null>(null)

  async function handleDelete(customer: AdminCustomer) {
    const confirmed = window.confirm(
      `Delete customer "${customer.name}" (#${customer.id})? This cannot be undone.`,
    )
    if (!confirmed) return

    setActionError(null)
    setActionSuccess(null)
    setDeletingId(customer.id)
    try {
      const result = await deleteCustomer(customer.id)
      setActionSuccess(result.message || 'Customer deleted successfully')
      if (selected?.id === customer.id) {
        setSelected(null)
      }
      await refresh()
    } catch (err) {
      setActionError(getErrorMessage(err, 'Failed to delete customer'))
    } finally {
      setDeletingId(null)
    }
  }

  return (
    <div>
      <PageHeader
        title="Customers"
        description="View, update, and delete customer accounts."
        actions={
          <Button variant="secondary" size="sm" onClick={() => void refresh()} disabled={loading}>
            Refresh
          </Button>
        }
      />

      {error ? <Alert variant="error">{error}</Alert> : null}
      {actionError ? <Alert variant="error">{actionError}</Alert> : null}
      {actionSuccess ? <Alert variant="success">{actionSuccess}</Alert> : null}

      {selected ? (
        <div className="mb-6">
          <CustomerEditForm
            key={`${selected.id}-${selected.name}-${selected.email}`}
            customer={selected}
            onCancel={() => setSelected(null)}
            onSuccess={() => {
              void refresh()
            }}
          />
        </div>
      ) : null}

      {loading ? (
        <Spinner label="Loading customers…" />
      ) : (
        <CustomersTable
          customers={customers}
          onEdit={setSelected}
          onDelete={(customer) => void handleDelete(customer)}
          deletingId={deletingId}
        />
      )}
    </div>
  )
}

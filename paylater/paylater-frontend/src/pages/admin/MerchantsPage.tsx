import { useState } from 'react'
import { MerchantEditForm } from '../../components/admin/MerchantEditForm'
import { MerchantsTable } from '../../components/admin/MerchantsTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useAdminMerchants } from '../../hooks/useAdminMerchants'
import { deleteMerchant } from '../../services/api/admin.api'
import type { AdminMerchant } from '../../types/admin'
import { getErrorMessage } from '../../types/api'

export function AdminMerchantsPage() {
  const { merchants, loading, error, refresh } = useAdminMerchants()
  const [selected, setSelected] = useState<AdminMerchant | null>(null)
  const [actionError, setActionError] = useState<string | null>(null)
  const [actionSuccess, setActionSuccess] = useState<string | null>(null)
  const [deletingId, setDeletingId] = useState<number | null>(null)

  async function handleDelete(merchant: AdminMerchant) {
    const confirmed = window.confirm(
      `Delete merchant "${merchant.name}" (#${merchant.id})? This cannot be undone.`,
    )
    if (!confirmed) return

    setActionError(null)
    setActionSuccess(null)
    setDeletingId(merchant.id)
    try {
      const result = await deleteMerchant(merchant.id)
      setActionSuccess(result.message || 'Merchant deleted successfully')
      if (selected?.id === merchant.id) {
        setSelected(null)
      }
      await refresh()
    } catch (err) {
      setActionError(getErrorMessage(err, 'Failed to delete merchant'))
    } finally {
      setDeletingId(null)
    }
  }

  return (
    <div>
      <PageHeader
        title="Merchants"
        description="Manage merchant profiles and commission rates."
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
          <MerchantEditForm
            key={`${selected.id}-${selected.name}-${selected.email}-${selected.commission_percentage}`}
            merchant={selected}
            onCancel={() => setSelected(null)}
            onSuccess={() => {
              void refresh()
            }}
          />
        </div>
      ) : null}

      {loading ? (
        <Spinner label="Loading merchants…" />
      ) : (
        <MerchantsTable
          merchants={merchants}
          onEdit={setSelected}
          onDelete={(merchant) => void handleDelete(merchant)}
          deletingId={deletingId}
        />
      )}
    </div>
  )
}

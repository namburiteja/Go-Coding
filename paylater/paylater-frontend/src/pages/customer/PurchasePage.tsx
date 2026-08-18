import { PurchaseForm } from '../../components/customer/PurchaseForm'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useCustomerProfile } from '../../hooks/useCustomerProfile'

export function CustomerPurchasePage() {
  const { profile, loading, error, refresh } = useCustomerProfile()

  if (loading) {
    return <Spinner label="Loading purchase…" />
  }

  if (error || !profile) {
    return (
      <div className="space-y-3">
        <Alert variant="error">{error || 'Unable to load account'}</Alert>
        <Button variant="secondary" onClick={() => void refresh()}>
          Retry
        </Button>
      </div>
    )
  }

  return (
    <div>
      <PageHeader
        title="Purchase with PayLater"
        description="Buy now and settle on your payment due date."
      />
      <PurchaseForm profile={profile} onSuccess={() => void refresh()} />
    </div>
  )
}

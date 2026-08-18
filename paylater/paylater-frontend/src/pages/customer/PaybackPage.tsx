import { PaybackForm } from '../../components/customer/PaybackForm'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useCustomerProfile } from '../../hooks/useCustomerProfile'

export function CustomerPaybackPage() {
  const { profile, loading, error, refresh } = useCustomerProfile()

  if (loading) {
    return <Spinner label="Loading payback…" />
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
        title="Payback"
        description="Clear part or all of your outstanding PayLater balance."
      />
      <PaybackForm profile={profile} onSuccess={() => void refresh()} />
    </div>
  )
}

import { MerchantProfileForm } from '../../components/merchant/MerchantProfileForm'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useMerchantProfile } from '../../hooks/useMerchantProfile'

export function MerchantProfilePage() {
  const { profile, loading, error, refresh } = useMerchantProfile()

  if (loading) {
    return <Spinner label="Loading profile…" />
  }

  if (error || !profile) {
    return (
      <div className="space-y-3">
        <Alert variant="error">{error || 'Unable to load profile'}</Alert>
        <Button variant="secondary" onClick={() => void refresh()}>
          Retry
        </Button>
      </div>
    )
  }

  return (
    <div>
      <PageHeader
        title="Profile"
        description="View and update your merchant account details."
      />
      <MerchantProfileForm
        key={`${profile.id}-${profile.name}-${profile.email}`}
        profile={profile}
        onSuccess={() => void refresh()}
      />
    </div>
  )
}

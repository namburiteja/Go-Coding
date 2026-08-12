import { AdminProfileForm } from '../../components/admin/AdminProfileForm'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useAdminProfile } from '../../hooks/useAdminProfile'

export function AdminProfilePage() {
  const { profile, loading, error, refresh } = useAdminProfile()

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
        description="View and update your administrator account. Passwords and hashes are never displayed."
      />
      <AdminProfileForm
        key={`${profile.id}-${profile.name}-${profile.email}`}
        profile={profile}
        onSuccess={() => void refresh()}
      />
    </div>
  )
}

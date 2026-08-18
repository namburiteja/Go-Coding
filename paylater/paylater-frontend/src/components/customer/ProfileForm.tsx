import { type FormEvent, useState } from 'react'
import type { CustomerProfile } from '../../types/customer'
import { getErrorMessage } from '../../types/api'
import { formatDate } from '../../utils/credit'
import { updateCustomerProfile } from '../../services/api/customer.api'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { Input } from '../ui/Input'
import { StatusBadge } from '../ui/StatusBadge'

interface ProfileFormProps {
  profile: CustomerProfile
  onSuccess?: () => void
}

export function ProfileForm({ profile, onSuccess }: ProfileFormProps) {
  const [name, setName] = useState(profile.name)
  const [email, setEmail] = useState(profile.email)
  const [success, setSuccess] = useState<string | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  async function onSubmit(event: FormEvent) {
    event.preventDefault()
    setError(null)
    setSuccess(null)

    if (!name.trim() || !email.trim()) {
      setError('Name and email are required.')
      return
    }

    setLoading(true)
    try {
      const result = await updateCustomerProfile({
        name: name.trim(),
        email: email.trim(),
      })
      setSuccess(result.message || 'Profile updated successfully')
      onSuccess?.()
    } catch (err) {
      setError(getErrorMessage(err, 'Failed to update profile'))
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="grid gap-4 lg:grid-cols-5">
      <Card className="lg:col-span-2">
        <p className="text-xs uppercase tracking-wide text-slate-500">Account</p>
        <h3 className="mt-2 text-lg font-semibold text-slate-900">{profile.name}</h3>
        <p className="mt-1 text-sm text-slate-600">{profile.email}</p>
        <div className="mt-4 space-y-2 text-sm text-slate-600">
          <p className="flex items-center gap-2">
            Status: <StatusBadge status={profile.status ?? 'ACTIVE'} />
          </p>
          <p>
            Customer ID: <span className="font-medium text-slate-900">#{profile.id}</span>
          </p>
          <p>
            Member since:{' '}
            <span className="font-medium text-slate-900">
              {formatDate(profile.created_at)}
            </span>
          </p>
        </div>
      </Card>

      <Card className="lg:col-span-3">
        <h3 className="text-base font-semibold text-slate-900">Edit profile</h3>
        <p className="mt-1 text-sm text-slate-600">Update your display name and email.</p>
        <form className="mt-4 space-y-4" onSubmit={onSubmit}>
          {error ? <Alert variant="error">{error}</Alert> : null}
          {success ? <Alert variant="success">{success}</Alert> : null}
          <Input
            label="Full name"
            name="name"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <Input
            label="Email"
            type="email"
            name="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <Button type="submit" loading={loading}>
            Save changes
          </Button>
        </form>
      </Card>
    </div>
  )
}

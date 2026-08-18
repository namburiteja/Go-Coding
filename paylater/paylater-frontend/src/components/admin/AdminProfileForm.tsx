import { type FormEvent, useState } from 'react'
import type { AdminProfile } from '../../types/admin'
import { getErrorMessage } from '../../types/api'
import { updateAdmin } from '../../services/api/admin.api'
import { formatDate } from '../../utils/credit'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { Input } from '../ui/Input'

interface AdminProfileFormProps {
  profile: AdminProfile
  title?: string
  description?: string
  onCancel?: () => void
  onSuccess?: () => void
}

export function AdminProfileForm({
  profile,
  title = 'Edit profile',
  description = 'Update your display name and email. Passwords are never shown.',
  onCancel,
  onSuccess,
}: AdminProfileFormProps) {
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
      const result = await updateAdmin(profile.id, {
        name: name.trim(),
        email: email.trim(),
      })
      setSuccess(result.message || 'Admin updated successfully')
      onSuccess?.()
    } catch (err) {
      setError(getErrorMessage(err, 'Failed to update admin'))
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
          <p>
            Admin ID: <span className="font-medium text-slate-900">#{profile.id}</span>
          </p>
          <p>
            Role: <span className="font-medium text-slate-900">{profile.role}</span>
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
        <div className="flex items-start justify-between gap-3">
          <div>
            <h3 className="text-base font-semibold text-slate-900">{title}</h3>
            <p className="mt-1 text-sm text-slate-600">{description}</p>
          </div>
          {onCancel ? (
            <Button variant="ghost" size="sm" onClick={onCancel}>
              Close
            </Button>
          ) : null}
        </div>
        <form className="mt-4 space-y-4" onSubmit={onSubmit}>
          {error ? <Alert variant="error">{error}</Alert> : null}
          {success ? <Alert variant="success">{success}</Alert> : null}
          <Input
            label="Name"
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
          <div className="flex flex-wrap gap-2">
            <Button type="submit" loading={loading}>
              Save changes
            </Button>
            {onCancel ? (
              <Button variant="secondary" onClick={onCancel}>
                Cancel
              </Button>
            ) : null}
          </div>
        </form>
      </Card>
    </div>
  )
}

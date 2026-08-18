import { type FormEvent, useState } from 'react'
import type { AdminMerchant } from '../../types/admin'
import { getErrorMessage } from '../../types/api'
import { updateMerchant, updateMerchantCommission } from '../../services/api/admin.api'
import { formatDate } from '../../utils/credit'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { Input } from '../ui/Input'

interface MerchantEditFormProps {
  merchant: AdminMerchant
  onCancel: () => void
  onSuccess: () => void
}

export function MerchantEditForm({ merchant, onCancel, onSuccess }: MerchantEditFormProps) {
  const [name, setName] = useState(merchant.name)
  const [email, setEmail] = useState(merchant.email)
  const [commission, setCommission] = useState(merchant.commission_percentage ?? '')
  const [success, setSuccess] = useState<string | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [savingProfile, setSavingProfile] = useState(false)
  const [savingCommission, setSavingCommission] = useState(false)

  async function onSubmitProfile(event: FormEvent) {
    event.preventDefault()
    setError(null)
    setSuccess(null)

    if (!name.trim() || !email.trim()) {
      setError('Name and email are required.')
      return
    }

    setSavingProfile(true)
    try {
      const result = await updateMerchant(merchant.id, {
        name: name.trim(),
        email: email.trim(),
      })
      setSuccess(result.message || 'Merchant updated successfully')
      onSuccess()
    } catch (err) {
      setError(getErrorMessage(err, 'Failed to update merchant'))
    } finally {
      setSavingProfile(false)
    }
  }

  async function onSubmitCommission(event: FormEvent) {
    event.preventDefault()
    setError(null)
    setSuccess(null)

    if (!commission.trim()) {
      setError('Commission percentage is required.')
      return
    }

    setSavingCommission(true)
    try {
      const result = await updateMerchantCommission(merchant.id, {
        commission_percentage: commission.trim(),
      })
      setSuccess(result.message || 'Commission updated successfully')
      onSuccess()
    } catch (err) {
      setError(getErrorMessage(err, 'Failed to update commission'))
    } finally {
      setSavingCommission(false)
    }
  }

  return (
    <Card>
      <div className="mb-4 flex flex-col gap-2 sm:flex-row sm:items-start sm:justify-between">
        <div>
          <h3 className="text-base font-semibold text-slate-900">Edit merchant</h3>
          <p className="mt-1 text-sm text-slate-600">
            Update profile details or commission rate using the dedicated admin APIs.
          </p>
        </div>
        <Button variant="ghost" size="sm" onClick={onCancel}>
          Close
        </Button>
      </div>

      <div className="mb-4 grid gap-2 rounded-xl bg-slate-50 p-3 text-sm text-slate-600 sm:grid-cols-2">
        <p>
          ID: <span className="font-medium text-slate-900">#{merchant.id}</span>
        </p>
        <p>
          Phone: <span className="font-medium text-slate-900">{merchant.phone}</span>
        </p>
        <p>
          Current commission:{' '}
          <span className="font-medium text-slate-900">
            {merchant.commission_percentage != null
              ? `${merchant.commission_percentage}%`
              : '—'}
          </span>
        </p>
        <p>
          Created:{' '}
          <span className="font-medium text-slate-900">
            {formatDate(merchant.created_at)}
          </span>
        </p>
      </div>

      {error ? <Alert variant="error">{error}</Alert> : null}
      {success ? (
        <div className="mb-4">
          <Alert variant="success">{success}</Alert>
        </div>
      ) : null}

      <form className="space-y-4 border-b border-slate-100 pb-6" onSubmit={onSubmitProfile}>
        <h4 className="text-sm font-semibold text-slate-900">Profile</h4>
        <Input
          label="Business name"
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
        <Input label="Phone" name="phone" value={merchant.phone} disabled readOnly />
        <Button type="submit" loading={savingProfile}>
          Save profile
        </Button>
      </form>

      <form className="mt-6 space-y-4" onSubmit={onSubmitCommission}>
        <h4 className="text-sm font-semibold text-slate-900">Commission</h4>
        <Input
          label="Commission percentage"
          name="commission_percentage"
          required
          value={commission}
          onChange={(e) => setCommission(e.target.value)}
          placeholder="e.g. 2.5"
        />
        <div className="flex flex-wrap gap-2">
          <Button type="submit" loading={savingCommission}>
            Update commission
          </Button>
          <Button variant="secondary" onClick={onCancel}>
            Cancel
          </Button>
        </div>
      </form>
    </Card>
  )
}

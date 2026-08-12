import { type FormEvent, useState } from 'react'
import type { AdminCustomer } from '../../types/admin'
import { getErrorMessage } from '../../types/api'
import { updateCustomer } from '../../services/api/admin.api'
import { formatCurrency, formatDate } from '../../utils/credit'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { Input } from '../ui/Input'
import { StatusBadge } from '../ui/StatusBadge'

interface CustomerEditFormProps {
  customer: AdminCustomer
  onCancel: () => void
  onSuccess: () => void
}

export function CustomerEditForm({ customer, onCancel, onSuccess }: CustomerEditFormProps) {
  const [name, setName] = useState(customer.name)
  const [email, setEmail] = useState(customer.email)
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
      const result = await updateCustomer(customer.id, {
        name: name.trim(),
        email: email.trim(),
      })
      setSuccess(result.message || 'Customer updated successfully')
      onSuccess()
    } catch (err) {
      setError(getErrorMessage(err, 'Failed to update customer'))
    } finally {
      setLoading(false)
    }
  }

  return (
    <Card>
      <div className="mb-4 flex flex-col gap-2 sm:flex-row sm:items-start sm:justify-between">
        <div>
          <h3 className="text-base font-semibold text-slate-900">Edit customer</h3>
          <p className="mt-1 text-sm text-slate-600">
            Update name and email only. Credit limit, due, and status are system-managed.
          </p>
        </div>
        <Button variant="ghost" size="sm" onClick={onCancel}>
          Close
        </Button>
      </div>

      <div className="mb-4 grid gap-2 rounded-xl bg-slate-50 p-3 text-sm text-slate-600 sm:grid-cols-2">
        <p>
          ID: <span className="font-medium text-slate-900">#{customer.id}</span>
        </p>
        <p className="flex items-center gap-2">
          Status:{' '}
          {customer.status ? <StatusBadge status={customer.status} /> : '—'}
        </p>
        <p>
          Credit limit:{' '}
          <span className="font-medium text-slate-900">
            {formatCurrency(customer.credit_limit)}
          </span>
        </p>
        <p>
          Total due:{' '}
          <span className="font-medium text-slate-900">
            {formatCurrency(customer.total_due)}
          </span>
        </p>
        <p>
          Payment due:{' '}
          <span className="font-medium text-slate-900">
            {formatDate(customer.payment_due_date)}
          </span>
        </p>
        <p>
          Created:{' '}
          <span className="font-medium text-slate-900">
            {formatDate(customer.created_at)}
          </span>
        </p>
      </div>

      <form className="space-y-4" onSubmit={onSubmit}>
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
          <Button variant="secondary" onClick={onCancel}>
            Cancel
          </Button>
        </div>
      </form>
    </Card>
  )
}

import { type FormEvent, useState } from 'react'
import type { CustomerProfile } from '../../types/customer'
import { getErrorMessage } from '../../types/api'
import { formatCurrency, parseAmount } from '../../utils/credit'
import { payback } from '../../services/api/customer.api'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { Input } from '../ui/Input'

interface PaybackFormProps {
  profile: CustomerProfile
  onSuccess?: () => void
}

export function PaybackForm({ profile, onSuccess }: PaybackFormProps) {
  const [amount, setAmount] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  const due = parseAmount(profile.total_due)

  async function onSubmit(event: FormEvent) {
    event.preventDefault()
    setError(null)
    setSuccess(null)

    const value = parseAmount(amount)
    if (due <= 0) {
      setError('You have no outstanding due.')
      return
    }
    if (value <= 0) {
      setError('Amount must be greater than zero.')
      return
    }
    if (value > due) {
      setError(`Payback cannot exceed current due of ${formatCurrency(due)}.`)
      return
    }

    setLoading(true)
    try {
      const result = await payback({ amount: value.toFixed(2) })
      setSuccess(result.message || 'Payment successful')
      setAmount('')
      onSuccess?.()
    } catch (err) {
      setError(getErrorMessage(err, 'Payback failed'))
    } finally {
      setLoading(false)
    }
  }

  return (
    <Card className="max-w-xl">
      <form className="space-y-4" onSubmit={onSubmit}>
        {error ? <Alert variant="error">{error}</Alert> : null}
        {success ? <Alert variant="success">{success}</Alert> : null}

        <div className="rounded-xl bg-slate-50 px-4 py-3 text-sm text-slate-600">
          <p>
            Outstanding due:{' '}
            <span className="font-semibold text-slate-900">{formatCurrency(due)}</span>
          </p>
          <p className="mt-1 text-xs text-slate-500">
            Blocked customers can still pay their outstanding balance.
          </p>
        </div>

        <Input
          label="Payback amount"
          name="amount"
          type="number"
          min={0.01}
          step="0.01"
          required
          placeholder="e.g. 100.00"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          disabled={due <= 0}
        />

        <div className="flex flex-wrap gap-2">
          <Button
            type="button"
            variant="secondary"
            size="sm"
            disabled={due <= 0}
            onClick={() => setAmount(due.toFixed(2))}
          >
            Pay full due
          </Button>
          <Button type="submit" loading={loading} disabled={due <= 0}>
            Confirm payment
          </Button>
        </div>
      </form>
    </Card>
  )
}

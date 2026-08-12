import { type FormEvent, useEffect, useState } from 'react'
import type { CustomerProfile, MerchantOption } from '../../types/customer'
import { getErrorMessage } from '../../types/api'
import {
  availableCredit,
  formatCurrency,
  parseAmount,
} from '../../utils/credit'
import { getMerchantOptions, purchase } from '../../services/api/customer.api'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { Input } from '../ui/Input'
import { Select } from '../ui/Select'
import { Spinner } from '../ui/Spinner'

interface PurchaseFormProps {
  profile: CustomerProfile
  onSuccess?: () => void
}

export function PurchaseForm({ profile, onSuccess }: PurchaseFormProps) {
  const [merchants, setMerchants] = useState<MerchantOption[]>([])
  const [merchantsLoading, setMerchantsLoading] = useState(true)
  const [merchantsError, setMerchantsError] = useState<string | null>(null)
  const [merchantId, setMerchantId] = useState('')
  const [amount, setAmount] = useState('')
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  const due = profile.total_due ?? '0'
  const available = availableCredit(profile.credit_limit, due)
  const status = (profile.status ?? 'ACTIVE').toUpperCase()

  useEffect(() => {
    let cancelled = false

    async function loadMerchants() {
      try {
        const options = await getMerchantOptions()
        if (!cancelled) {
          setMerchants(options)
          setMerchantsError(null)
        }
      } catch (err) {
        if (!cancelled) {
          setMerchantsError(getErrorMessage(err, 'Failed to load merchants'))
        }
      } finally {
        if (!cancelled) {
          setMerchantsLoading(false)
        }
      }
    }

    void loadMerchants()
    return () => {
      cancelled = true
    }
  }, [])

  async function onSubmit(event: FormEvent) {
    event.preventDefault()
    setError(null)
    setSuccess(null)

    const merchant = Number.parseInt(merchantId, 10)
    const value = parseAmount(amount)

    if (!Number.isFinite(merchant) || merchant <= 0) {
      setError('Select a merchant.')
      return
    }
    if (value <= 0) {
      setError('Amount must be greater than zero.')
      return
    }
    if (status === 'BLOCKED') {
      setError('Your account is blocked. You can pay outstanding dues, but cannot purchase.')
      return
    }
    if (value > available) {
      setError(`Amount exceeds available credit of ${formatCurrency(available)}.`)
      return
    }

    setLoading(true)
    try {
      const result = await purchase({
        merchantId: merchant,
        amount: value.toFixed(2),
      })
      setSuccess(result.message || 'Purchase successful')
      setAmount('')
      onSuccess?.()
    } catch (err) {
      setError(getErrorMessage(err, 'Purchase failed'))
    } finally {
      setLoading(false)
    }
  }

  return (
    <Card className="max-w-xl">
      <form className="space-y-4" onSubmit={onSubmit}>
        {error ? <Alert variant="error">{error}</Alert> : null}
        {success ? <Alert variant="success">{success}</Alert> : null}
        {merchantsError ? <Alert variant="error">{merchantsError}</Alert> : null}

        <div className="rounded-xl bg-slate-50 px-4 py-3 text-sm text-slate-600">
          <p>
            Available credit:{' '}
            <span className="font-semibold text-slate-900">{formatCurrency(available)}</span>
          </p>
          {status === 'BLOCKED' ? (
            <p className="mt-1 text-xs text-red-600">Purchases are disabled while blocked.</p>
          ) : null}
        </div>

        {merchantsLoading ? (
          <Spinner label="Loading merchants…" />
        ) : (
          <Select
            label="Merchant"
            name="merchantId"
            required
            value={merchantId}
            onChange={(e) => setMerchantId(e.target.value)}
            placeholder="Choose a merchant"
            options={merchants.map((m) => ({
              value: String(m.id),
              label: m.name,
            }))}
            disabled={merchants.length === 0 || status === 'BLOCKED'}
          />
        )}

        <Input
          label="Amount"
          name="amount"
          type="number"
          min={0.01}
          step="0.01"
          required
          placeholder="e.g. 100.00"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          disabled={status === 'BLOCKED'}
        />

        <Button
          type="submit"
          className="w-full sm:w-auto"
          loading={loading}
          disabled={status === 'BLOCKED' || merchantsLoading || merchants.length === 0}
        >
          Confirm purchase
        </Button>
      </form>
    </Card>
  )
}

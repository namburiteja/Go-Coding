import { Link } from 'react-router-dom'
import { useEffect, useMemo, useState } from 'react'
import { CreditOverview } from '../../components/customer/CreditOverview'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { TransactionTable } from '../../components/ui/TransactionTable'
import { useCustomerProfile } from '../../hooks/useCustomerProfile'
import {
  getCustomerTransactions,
  getMerchantOptions,
} from '../../services/api/customer.api'
import type { CustomerTransaction } from '../../types/customer'
import { getErrorMessage } from '../../types/api'

export function CustomerDashboardPage() {
  const { profile, loading, error, refresh } = useCustomerProfile()
  const [transactions, setTransactions] = useState<CustomerTransaction[]>([])
  const [merchantNames, setMerchantNames] = useState<Record<number, string>>({})
  const [txLoading, setTxLoading] = useState(true)
  const [txError, setTxError] = useState<string | null>(null)

  useEffect(() => {
    let cancelled = false

    async function load() {
      try {
        const [txs, merchants] = await Promise.all([
          getCustomerTransactions(),
          getMerchantOptions().catch(() => []),
        ])
        if (cancelled) return
        setTransactions(txs)
        setMerchantNames(Object.fromEntries(merchants.map((m) => [m.id, m.name])))
        setTxError(null)
      } catch (err) {
        if (!cancelled) {
          setTxError(getErrorMessage(err, 'Failed to load transactions'))
        }
      } finally {
        if (!cancelled) {
          setTxLoading(false)
        }
      }
    }

    void load()
    return () => {
      cancelled = true
    }
  }, [])

  const recent = useMemo(() => transactions.slice(0, 5), [transactions])

  if (loading) {
    return <Spinner label="Loading dashboard…" />
  }

  if (error || !profile) {
    return (
      <div className="space-y-3">
        <Alert variant="error">{error || 'Unable to load dashboard'}</Alert>
        <Button variant="secondary" onClick={() => void refresh()}>
          Retry
        </Button>
      </div>
    )
  }

  return (
    <div>
      <PageHeader
        title="Dashboard"
        description="Your PayLater credit overview and recent activity."
        actions={
          <>
            <Link to="/customer/purchase">
              <Button>New purchase</Button>
            </Link>
            <Link to="/customer/payback">
              <Button variant="secondary">Make payment</Button>
            </Link>
          </>
        }
      />

      <CreditOverview profile={profile} />

      <div className="mt-8">
        <div className="mb-3 flex items-center justify-between">
          <h3 className="text-base font-semibold text-slate-900">Recent transactions</h3>
          <Link to="/customer/transactions" className="text-sm font-medium text-slate-700 underline">
            View all
          </Link>
        </div>
        {txError ? <Alert variant="error">{txError}</Alert> : null}
        {txLoading ? (
          <Spinner label="Loading transactions…" />
        ) : (
          <TransactionTable transactions={recent} merchantNames={merchantNames} />
        )}
      </div>
    </div>
  )
}

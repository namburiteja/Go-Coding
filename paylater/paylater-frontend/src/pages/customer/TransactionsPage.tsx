import { useCallback, useEffect, useState } from 'react'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { TransactionTable } from '../../components/ui/TransactionTable'
import {
  getCustomerTransactions,
  getMerchantOptions,
} from '../../services/api/customer.api'
import type { CustomerTransaction } from '../../types/customer'
import { getErrorMessage } from '../../types/api'

export function CustomerTransactionsPage() {
  const [transactions, setTransactions] = useState<CustomerTransaction[]>([])
  const [merchantNames, setMerchantNames] = useState<Record<number, string>>({})
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refresh = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const [txs, merchants] = await Promise.all([
        getCustomerTransactions(),
        getMerchantOptions().catch(() => []),
      ])
      setTransactions(txs)
      setMerchantNames(Object.fromEntries(merchants.map((m) => [m.id, m.name])))
    } catch (err) {
      setError(getErrorMessage(err, 'Failed to load transactions'))
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    let cancelled = false

    async function loadInitial() {
      try {
        const [txs, merchants] = await Promise.all([
          getCustomerTransactions(),
          getMerchantOptions().catch(() => []),
        ])
        if (cancelled) return
        setTransactions(txs)
        setMerchantNames(Object.fromEntries(merchants.map((m) => [m.id, m.name])))
        setError(null)
      } catch (err) {
        if (!cancelled) {
          setError(getErrorMessage(err, 'Failed to load transactions'))
        }
      } finally {
        if (!cancelled) {
          setLoading(false)
        }
      }
    }

    void loadInitial()
    return () => {
      cancelled = true
    }
  }, [])

  return (
    <div>
      <PageHeader
        title="Transaction history"
        description="All of your PayLater purchases and payments."
        actions={
          <Button variant="secondary" size="sm" onClick={() => void refresh()} disabled={loading}>
            Refresh
          </Button>
        }
      />
      {error ? <Alert variant="error">{error}</Alert> : null}
      {loading ? (
        <Spinner label="Loading transactions…" />
      ) : (
        <TransactionTable transactions={transactions} merchantNames={merchantNames} />
      )}
    </div>
  )
}

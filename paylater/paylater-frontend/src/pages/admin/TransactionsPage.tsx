import { useMemo } from 'react'
import { AdminTransactionsTable } from '../../components/admin/AdminTransactionsTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { StatCard } from '../../components/ui/StatCard'
import { useAdminCustomers } from '../../hooks/useAdminCustomers'
import { useAdminMerchants } from '../../hooks/useAdminMerchants'
import { useAdminTransactions } from '../../hooks/useAdminTransactions'
import { formatCurrency, parseAmount } from '../../utils/credit'

export function AdminTransactionsPage() {
  const { transactions, loading, error, refresh } = useAdminTransactions()
  const { customers } = useAdminCustomers()
  const { merchants } = useAdminMerchants()

  const customerNames = useMemo(
    () => Object.fromEntries(customers.map((c) => [c.id, c.name])),
    [customers],
  )
  const merchantNames = useMemo(
    () => Object.fromEntries(merchants.map((m) => [m.id, m.name])),
    [merchants],
  )

  const summary = useMemo(() => {
    let purchases = 0
    let paybacks = 0
    let volume = 0
    let fees = 0
    for (const tx of transactions) {
      if (tx.transaction_type === 'PURCHASE') {
        purchases += 1
        volume += parseAmount(tx.amount)
        fees += parseAmount(tx.commission_amount)
      } else if (tx.transaction_type === 'PAYBACK') {
        paybacks += 1
      }
    }
    return { purchases, paybacks, volume, fees }
  }, [transactions])

  return (
    <div>
      <PageHeader
        title="Transactions"
        description="All ledger purchases and paybacks across the platform."
        actions={
          <Button variant="secondary" size="sm" onClick={() => void refresh()} disabled={loading}>
            Refresh
          </Button>
        }
      />

      {error ? <Alert variant="error">{error}</Alert> : null}

      <div className="mb-6 grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        <StatCard label="Purchases" value={String(summary.purchases)} />
        <StatCard label="Paybacks" value={String(summary.paybacks)} />
        <StatCard label="Purchase volume" value={formatCurrency(summary.volume)} />
        <StatCard
          label="Fees collected"
          value={formatCurrency(summary.fees)}
          tone="warning"
        />
      </div>

      {loading ? (
        <Spinner label="Loading transactions…" />
      ) : (
        <AdminTransactionsTable
          transactions={transactions}
          customerNames={customerNames}
          merchantNames={merchantNames}
        />
      )}
    </div>
  )
}

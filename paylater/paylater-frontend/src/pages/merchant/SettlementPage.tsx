import { useMemo } from 'react'
import { MerchantSalesTable } from '../../components/merchant/MerchantSalesTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { Card } from '../../components/ui/Card'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { StatCard } from '../../components/ui/StatCard'
import { useMerchantProfile } from '../../hooks/useMerchantProfile'
import { useMerchantTransactions } from '../../hooks/useMerchantTransactions'
import { formatCurrency } from '../../utils/credit'
import { summarizeMerchantSettlement } from '../../utils/merchantStats'

export function MerchantSettlementPage() {
  const {
    profile,
    loading: profileLoading,
    error: profileError,
    refresh: refreshProfile,
  } = useMerchantProfile()
  const {
    transactions,
    loading: txLoading,
    error: txError,
    refresh: refreshTx,
  } = useMerchantTransactions()

  const settlement = useMemo(
    () => summarizeMerchantSettlement(transactions),
    [transactions],
  )
  const sales = useMemo(
    () => transactions.filter((tx) => tx.transaction_type === 'PURCHASE'),
    [transactions],
  )

  const loading = profileLoading || txLoading
  const error = profileError || txError

  if (loading && !profile) {
    return <Spinner label="Loading settlement…" />
  }

  if (error && !profile) {
    return (
      <div className="space-y-3">
        <Alert variant="error">{error}</Alert>
        <Button
          variant="secondary"
          onClick={() => {
            void refreshProfile()
            void refreshTx()
          }}
        >
          Retry
        </Button>
      </div>
    )
  }

  return (
    <div>
      <PageHeader
        title="Settlement"
        description="Fees and net amounts derived from your sales transactions."
        actions={
          <Button
            variant="secondary"
            size="sm"
            onClick={() => {
              void refreshProfile()
              void refreshTx()
            }}
            disabled={loading}
          >
            Refresh
          </Button>
        }
      />

      <Card className="mb-4 border-dashed bg-slate-50">
        <p className="text-sm text-slate-600">
          Settlement is calculated from your commission rate and recorded sales (gross sales,
          fees, and net amount).
        </p>
      </Card>

      {error ? <Alert variant="error">{error}</Alert> : null}

      <div className="mb-6 grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        <StatCard
          label="Commission rate"
          value={
            profile?.commission_percentage != null
              ? `${profile.commission_percentage}%`
              : '—'
          }
          hint="Configured on your merchant account"
        />
        <StatCard
          label="Gross sales"
          value={formatCurrency(settlement.totalSales)}
          hint={`${settlement.saleCount} purchases`}
        />
        <StatCard
          label="Total fees"
          value={formatCurrency(settlement.totalCommission)}
          hint="Sum of commission_amount"
          tone="warning"
        />
        <StatCard
          label="Net to merchant"
          value={formatCurrency(settlement.netSettlement)}
          hint="Gross sales − fees"
          tone="success"
        />
      </div>

      <h3 className="mb-3 text-base font-semibold text-slate-900">Sales included in settlement</h3>
      {txLoading ? (
        <Spinner label="Loading sales…" />
      ) : (
        <MerchantSalesTable transactions={sales} />
      )}
    </div>
  )
}

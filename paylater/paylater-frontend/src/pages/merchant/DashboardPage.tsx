import { Link } from 'react-router-dom'
import { useMemo } from 'react'
import { MerchantOverview } from '../../components/merchant/MerchantOverview'
import { MerchantSalesTable } from '../../components/merchant/MerchantSalesTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useMerchantProfile } from '../../hooks/useMerchantProfile'
import { useMerchantTransactions } from '../../hooks/useMerchantTransactions'
import { summarizeMerchantSettlement } from '../../utils/merchantStats'

export function MerchantDashboardPage() {
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
  const recentSales = useMemo(
    () => transactions.filter((tx) => tx.transaction_type === 'PURCHASE').slice(0, 5),
    [transactions],
  )

  if (profileLoading) {
    return <Spinner label="Loading dashboard…" />
  }

  if (profileError || !profile) {
    return (
      <div className="space-y-3">
        <Alert variant="error">{profileError || 'Unable to load dashboard'}</Alert>
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
        title="Dashboard"
        description="Sales overview and settlement snapshot for your store."
        actions={
          <>
            <Link to="/merchant/transactions">
              <Button variant="secondary">View sales</Button>
            </Link>
            <Link to="/merchant/reports">
              <Button>Settlement</Button>
            </Link>
          </>
        }
      />

      <MerchantOverview profile={profile} settlement={settlement} />

      <div className="mt-8">
        <div className="mb-3 flex items-center justify-between">
          <h3 className="text-base font-semibold text-slate-900">Recent sales</h3>
          <Link
            to="/merchant/transactions"
            className="text-sm font-medium text-slate-700 underline"
          >
            View all
          </Link>
        </div>
        {txError ? <Alert variant="error">{txError}</Alert> : null}
        {txLoading ? (
          <Spinner label="Loading sales…" />
        ) : (
          <MerchantSalesTable transactions={recentSales} />
        )}
      </div>
    </div>
  )
}

import { useMemo } from 'react'
import { MerchantSalesTable } from '../../components/merchant/MerchantSalesTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useMerchantTransactions } from '../../hooks/useMerchantTransactions'

export function MerchantTransactionsPage() {
  const { transactions, loading, error, refresh } = useMerchantTransactions()

  const sales = useMemo(
    () => transactions.filter((tx) => tx.transaction_type === 'PURCHASE'),
    [transactions],
  )

  return (
    <div>
      <PageHeader
        title="Sales & transactions"
        description="PayLater purchases recorded at your merchant account."
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
        <MerchantSalesTable transactions={sales} />
      )}
    </div>
  )
}

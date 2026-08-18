import { useMemo } from 'react'
import { MerchantCustomersTable } from '../../components/merchant/MerchantCustomersTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { Card } from '../../components/ui/Card'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useMerchantTransactions } from '../../hooks/useMerchantTransactions'
import { deriveMerchantCustomers } from '../../utils/merchantStats'

export function MerchantCustomersPage() {
  const { transactions, loading, error, refresh } = useMerchantTransactions()

  const customers = useMemo(
    () => deriveMerchantCustomers(transactions),
    [transactions],
  )

  return (
    <div>
      <PageHeader
        title="Customers"
        description="Buyers derived from your PayLater sales history."
        actions={
          <Button variant="secondary" size="sm" onClick={() => void refresh()} disabled={loading}>
            Refresh
          </Button>
        }
      />
      <Card className="mb-4 border-dashed bg-slate-50">
        <p className="text-sm text-slate-600">
          Customer details are summarized from your PayLater sales. Full customer profiles are
          not exposed on merchant APIs.
        </p>
      </Card>
      {error ? <Alert variant="error">{error}</Alert> : null}
      {loading ? (
        <Spinner label="Loading customers…" />
      ) : (
        <MerchantCustomersTable customers={customers} />
      )}
    </div>
  )
}

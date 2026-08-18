import { Link } from 'react-router-dom'
import { useMemo } from 'react'
import { AdminOverview } from '../../components/admin/AdminOverview'
import { AdminTransactionsTable } from '../../components/admin/AdminTransactionsTable'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { PageHeader } from '../../components/ui/PageHeader'
import { Spinner } from '../../components/ui/Spinner'
import { useAdminAdmins } from '../../hooks/useAdminAdmins'
import { useAdminCustomers } from '../../hooks/useAdminCustomers'
import { useAdminMerchants } from '../../hooks/useAdminMerchants'
import { useAdminProfile } from '../../hooks/useAdminProfile'
import { useAdminTransactions } from '../../hooks/useAdminTransactions'
import { summarizeAdminDashboard } from '../../utils/adminStats'

export function AdminDashboardPage() {
  const {
    profile,
    loading: profileLoading,
    error: profileError,
    refresh: refreshProfile,
  } = useAdminProfile()
  const {
    customers,
    loading: customersLoading,
    error: customersError,
    refresh: refreshCustomers,
  } = useAdminCustomers()
  const {
    merchants,
    loading: merchantsLoading,
    error: merchantsError,
    refresh: refreshMerchants,
  } = useAdminMerchants()
  const {
    admins,
    loading: adminsLoading,
    error: adminsError,
    refresh: refreshAdmins,
  } = useAdminAdmins()
  const {
    transactions,
    loading: txLoading,
    error: txError,
    refresh: refreshTx,
  } = useAdminTransactions()

  const stats = useMemo(
    () =>
      summarizeAdminDashboard(
        customers,
        merchants.length,
        admins.length,
        transactions,
      ),
    [customers, merchants.length, admins.length, transactions],
  )

  const customerNames = useMemo(
    () => Object.fromEntries(customers.map((c) => [c.id, c.name])),
    [customers],
  )
  const merchantNames = useMemo(
    () => Object.fromEntries(merchants.map((m) => [m.id, m.name])),
    [merchants],
  )
  const recent = useMemo(() => transactions.slice(0, 8), [transactions])

  const loading =
    profileLoading || customersLoading || merchantsLoading || adminsLoading
  const listError = customersError || merchantsError || adminsError

  function refreshAll() {
    void refreshProfile()
    void refreshCustomers()
    void refreshMerchants()
    void refreshAdmins()
    void refreshTx()
  }

  if (loading && !profile) {
    return <Spinner label="Loading dashboard…" />
  }

  if ((profileError || !profile) && !profileLoading) {
    return (
      <div className="space-y-3">
        <Alert variant="error">{profileError || 'Unable to load dashboard'}</Alert>
        <Button variant="secondary" onClick={refreshAll}>
          Retry
        </Button>
      </div>
    )
  }

  if (!profile) {
    return <Spinner label="Loading dashboard…" />
  }

  return (
    <div>
      <PageHeader
        title="Dashboard"
        description="Platform overview across customers, merchants, and ledger activity."
        actions={
          <>
            <Button variant="secondary" size="sm" onClick={refreshAll} disabled={loading || txLoading}>
              Refresh
            </Button>
            <Link to="/admin/reports">
              <Button size="sm">Reports</Button>
            </Link>
          </>
        }
      />

      {listError ? <Alert variant="error">{listError}</Alert> : null}

      <AdminOverview profile={profile} stats={stats} />

      <div className="mt-8">
        <div className="mb-3 flex items-center justify-between">
          <h3 className="text-base font-semibold text-slate-900">Recent transactions</h3>
          <Link
            to="/admin/transactions"
            className="text-sm font-medium text-slate-700 underline"
          >
            View all
          </Link>
        </div>
        {txError ? <Alert variant="error">{txError}</Alert> : null}
        {txLoading ? (
          <Spinner label="Loading transactions…" />
        ) : (
          <AdminTransactionsTable
            transactions={recent}
            customerNames={customerNames}
            merchantNames={merchantNames}
          />
        )}
      </div>
    </div>
  )
}

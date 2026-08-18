import { type FormEvent, useState } from 'react'
import type { AdminCustomer, MerchantFeeReport } from '../../types/admin'
import { getErrorMessage } from '../../types/api'
import {
  getCustomerDueByName,
  getCustomersWithDue,
  getMerchantFees,
  getUsersAtCreditLimit,
} from '../../services/api/admin.api'
import { formatCurrency, formatDate } from '../../utils/credit'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { EmptyState } from '../ui/EmptyState'
import { Input } from '../ui/Input'
import { Spinner } from '../ui/Spinner'
import { StatusBadge } from '../ui/StatusBadge'

type ReportKind = 'credit-limit' | 'customers-due' | 'merchant-fees'

function CustomerReportTable({ customers }: { customers: AdminCustomer[] }) {
  if (customers.length === 0) {
    return <EmptyState title="No results" description="This report returned no customers." />
  }

  return (
    <div className="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-slate-200 text-left text-sm">
          <thead className="bg-slate-50 text-xs uppercase tracking-wide text-slate-500">
            <tr>
              <th className="px-4 py-3 font-medium">ID</th>
              <th className="px-4 py-3 font-medium">Name</th>
              <th className="px-4 py-3 font-medium">Email</th>
              <th className="px-4 py-3 font-medium">Credit limit</th>
              <th className="px-4 py-3 font-medium">Due</th>
              <th className="px-4 py-3 font-medium">Status</th>
              <th className="px-4 py-3 font-medium">Due date</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {customers.map((customer) => (
              <tr key={customer.id} className="hover:bg-slate-50/80">
                <td className="px-4 py-3 text-slate-700">#{customer.id}</td>
                <td className="px-4 py-3 font-medium text-slate-900">{customer.name}</td>
                <td className="px-4 py-3 text-slate-700">{customer.email}</td>
                <td className="px-4 py-3 text-slate-700">
                  {formatCurrency(customer.credit_limit)}
                </td>
                <td className="px-4 py-3 text-slate-700">
                  {formatCurrency(customer.total_due)}
                </td>
                <td className="px-4 py-3">
                  {customer.status ? <StatusBadge status={customer.status} /> : '—'}
                </td>
                <td className="px-4 py-3 text-slate-600">
                  {formatDate(customer.payment_due_date)}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

function MerchantFeesTable({ fees }: { fees: MerchantFeeReport[] }) {
  if (fees.length === 0) {
    return <EmptyState title="No merchant fees" description="No fee totals available yet." />
  }

  return (
    <div className="overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm">
      <div className="overflow-x-auto">
        <table className="min-w-full divide-y divide-slate-200 text-left text-sm">
          <thead className="bg-slate-50 text-xs uppercase tracking-wide text-slate-500">
            <tr>
              <th className="px-4 py-3 font-medium">ID</th>
              <th className="px-4 py-3 font-medium">Merchant</th>
              <th className="px-4 py-3 font-medium">Total fee collected</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {fees.map((row) => (
              <tr key={row.id} className="hover:bg-slate-50/80">
                <td className="px-4 py-3 text-slate-700">#{row.id}</td>
                <td className="px-4 py-3 font-medium text-slate-900">{row.name}</td>
                <td className="px-4 py-3 text-slate-900">
                  {formatCurrency(row.total_fee_collected)}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

export function ReportsPanel() {
  const [active, setActive] = useState<ReportKind | null>(null)
  const [customers, setCustomers] = useState<AdminCustomer[]>([])
  const [fees, setFees] = useState<MerchantFeeReport[]>([])
  const [lookupName, setLookupName] = useState('')
  const [lookupResult, setLookupResult] = useState<AdminCustomer | null>(null)
  const [loading, setLoading] = useState(false)
  const [lookupLoading, setLookupLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [lookupError, setLookupError] = useState<string | null>(null)

  async function loadReport(kind: ReportKind) {
    setActive(kind)
    setLoading(true)
    setError(null)
    setLookupResult(null)
    setLookupError(null)
    try {
      if (kind === 'credit-limit') {
        setCustomers(await getUsersAtCreditLimit())
        setFees([])
      } else if (kind === 'customers-due') {
        setCustomers(await getCustomersWithDue())
        setFees([])
      } else {
        setFees(await getMerchantFees())
        setCustomers([])
      }
    } catch (err) {
      setCustomers([])
      setFees([])
      setError(getErrorMessage(err, 'Failed to load report'))
    } finally {
      setLoading(false)
    }
  }

  async function onLookup(event: FormEvent) {
    event.preventDefault()
    setLookupError(null)
    setLookupResult(null)
    setError(null)
    setActive(null)

    const name = lookupName.trim()
    if (!name) {
      setLookupError('Enter a customer name.')
      return
    }

    setLookupLoading(true)
    try {
      const result = await getCustomerDueByName(name)
      setLookupResult(result)
    } catch (err) {
      setLookupError(getErrorMessage(err, 'Customer not found'))
    } finally {
      setLookupLoading(false)
    }
  }

  return (
    <div className="space-y-6">
      <Card>
        <h3 className="text-base font-semibold text-slate-900">Platform reports</h3>
        <p className="mt-1 text-sm text-slate-600">
          Uses existing admin report endpoints. Choose a report or look up a customer by exact name.
        </p>
        <div className="mt-4 flex flex-wrap gap-2">
          <Button
            variant={active === 'credit-limit' ? 'primary' : 'secondary'}
            size="sm"
            onClick={() => void loadReport('credit-limit')}
            disabled={loading}
          >
            At credit limit
          </Button>
          <Button
            variant={active === 'customers-due' ? 'primary' : 'secondary'}
            size="sm"
            onClick={() => void loadReport('customers-due')}
            disabled={loading}
          >
            Customers with due
          </Button>
          <Button
            variant={active === 'merchant-fees' ? 'primary' : 'secondary'}
            size="sm"
            onClick={() => void loadReport('merchant-fees')}
            disabled={loading}
          >
            Merchant fees
          </Button>
        </div>
      </Card>

      <Card>
        <h3 className="text-base font-semibold text-slate-900">Customer due by name</h3>
        <p className="mt-1 text-sm text-slate-600">
          Calls <code className="text-xs">GET /reports/customer-due/:name</code>.
        </p>
        <form className="mt-4 flex flex-col gap-3 sm:flex-row sm:items-end" onSubmit={onLookup}>
          <div className="flex-1">
            <Input
              label="Exact customer name"
              name="customer-name"
              value={lookupName}
              onChange={(e) => setLookupName(e.target.value)}
              placeholder="Jane Doe"
            />
          </div>
          <Button type="submit" loading={lookupLoading}>
            Look up
          </Button>
        </form>
        {lookupError ? (
          <div className="mt-3">
            <Alert variant="error">{lookupError}</Alert>
          </div>
        ) : null}
        {lookupResult ? (
          <div className="mt-4">
            <CustomerReportTable customers={[lookupResult]} />
          </div>
        ) : null}
      </Card>

      {error ? <Alert variant="error">{error}</Alert> : null}
      {loading ? <Spinner label="Loading report…" /> : null}

      {!loading && active === 'merchant-fees' ? <MerchantFeesTable fees={fees} /> : null}
      {!loading && (active === 'credit-limit' || active === 'customers-due') ? (
        <CustomerReportTable customers={customers} />
      ) : null}
    </div>
  )
}

import type { AdminDashboardStats, AdminProfile } from '../../types/admin'
import { formatCurrency, formatDate } from '../../utils/credit'
import { Card } from '../ui/Card'
import { StatCard } from '../ui/StatCard'

interface AdminOverviewProps {
  profile: AdminProfile
  stats: AdminDashboardStats
}

export function AdminOverview({ profile, stats }: AdminOverviewProps) {
  return (
    <div className="space-y-4">
      <Card className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <p className="text-sm text-slate-500">Administrator</p>
          <h3 className="mt-2 text-lg font-semibold text-slate-900">{profile.name}</h3>
          <p className="mt-1 text-sm text-slate-600">{profile.email}</p>
        </div>
        <div className="rounded-xl bg-slate-50 px-4 py-3 text-sm text-slate-600">
          <p>
            Role: <span className="font-semibold text-slate-900">{profile.role}</span>
          </p>
          <p className="mt-1 text-xs text-slate-500">
            Member since {formatDate(profile.created_at)}
          </p>
        </div>
      </Card>

      <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        <StatCard label="Customers" value={String(stats.customerCount)} hint="Registered accounts" />
        <StatCard label="Merchants" value={String(stats.merchantCount)} hint="Active stores" />
        <StatCard label="Admins" value={String(stats.adminCount)} hint="Platform operators" />
        <StatCard
          label="Transactions"
          value={String(stats.transactionCount)}
          hint={`${stats.purchaseCount} purchases · ${stats.paybackCount} paybacks`}
        />
      </div>

      <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        <StatCard
          label="Purchase volume"
          value={formatCurrency(stats.totalPurchaseVolume)}
          hint="Sum of PURCHASE amounts"
        />
        <StatCard
          label="Platform fees"
          value={formatCurrency(stats.totalFees)}
          hint="Sum of commission_amount"
          tone="warning"
        />
        <StatCard
          label="With outstanding due"
          value={String(stats.customersWithDue)}
          hint="Customers with total_due > 0"
        />
        <StatCard
          label="Blocked"
          value={String(stats.blockedCustomers)}
          hint="Customers with BLOCKED status"
          tone={stats.blockedCustomers > 0 ? 'danger' : 'default'}
        />
      </div>
    </div>
  )
}

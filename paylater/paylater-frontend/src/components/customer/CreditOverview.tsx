import type { CustomerProfile } from '../../types/customer'
import {
  availableCredit,
  formatCurrency,
  formatDate,
} from '../../utils/credit'
import { Card } from '../ui/Card'
import { StatCard } from '../ui/StatCard'
import { StatusBadge } from '../ui/StatusBadge'

interface CreditOverviewProps {
  profile: CustomerProfile
}

export function CreditOverview({ profile }: CreditOverviewProps) {
  const due = profile.total_due ?? '0'
  const available = availableCredit(profile.credit_limit, due)
  const dueAmount = Number.parseFloat(due)
  const dueTone = dueAmount > 0 ? 'warning' : 'success'

  return (
    <div className="space-y-4">
      <Card className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <p className="text-sm text-slate-500">Account status</p>
          <div className="mt-2 flex flex-wrap items-center gap-3">
            <h3 className="text-lg font-semibold text-slate-900">{profile.name}</h3>
            <StatusBadge status={profile.status ?? 'ACTIVE'} />
          </div>
          <p className="mt-1 text-sm text-slate-600">{profile.email}</p>
        </div>
        <div className="rounded-xl bg-slate-50 px-4 py-3 text-sm text-slate-600">
          <p>
            Payment due date:{' '}
            <span className="font-semibold text-slate-900">
              {formatDate(profile.payment_due_date)}
            </span>
          </p>
        </div>
      </Card>

      <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        <StatCard
          label="Credit limit"
          value={formatCurrency(profile.credit_limit)}
          hint="Maximum PayLater credit"
        />
        <StatCard
          label="Available credit"
          value={formatCurrency(available)}
          hint="Limit minus outstanding due"
          tone="success"
        />
        <StatCard
          label="Total due"
          value={formatCurrency(due)}
          hint="Outstanding balance"
          tone={dueTone}
        />
        <StatCard
          label="Payment due date"
          value={formatDate(profile.payment_due_date)}
          hint="Next applicable 5th"
        />
      </div>
    </div>
  )
}

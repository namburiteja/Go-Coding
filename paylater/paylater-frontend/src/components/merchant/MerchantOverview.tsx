import type { MerchantProfile, MerchantSettlementSummary } from '../../types/merchant'
import { formatCurrency, formatDate } from '../../utils/credit'
import { Card } from '../ui/Card'
import { StatCard } from '../ui/StatCard'

interface MerchantOverviewProps {
  profile: MerchantProfile
  settlement: MerchantSettlementSummary
}

export function MerchantOverview({ profile, settlement }: MerchantOverviewProps) {
  const commission = profile.commission_percentage
    ? `${profile.commission_percentage}%`
    : '—'

  return (
    <div className="space-y-4">
      <Card className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <p className="text-sm text-slate-500">Merchant account</p>
          <h3 className="mt-2 text-lg font-semibold text-slate-900">{profile.name}</h3>
          <p className="mt-1 text-sm text-slate-600">{profile.email}</p>
          <p className="mt-1 text-sm text-slate-600">Phone: {profile.phone}</p>
        </div>
        <div className="rounded-xl bg-slate-50 px-4 py-3 text-sm text-slate-600">
          <p>
            Commission rate:{' '}
            <span className="font-semibold text-slate-900">{commission}</span>
          </p>
          <p className="mt-1 text-xs text-slate-500">
            Member since {formatDate(profile.created_at)}
          </p>
        </div>
      </Card>

      <div className="grid gap-4 sm:grid-cols-2 xl:grid-cols-4">
        <StatCard
          label="Total sales"
          value={formatCurrency(settlement.totalSales)}
          hint={`${settlement.saleCount} PayLater purchases`}
        />
        <StatCard
          label="Commission fees"
          value={formatCurrency(settlement.totalCommission)}
          hint="Platform fee on sales"
          tone="warning"
        />
        <StatCard
          label="Net settlement"
          value={formatCurrency(settlement.netSettlement)}
          hint="Sales minus commission"
          tone="success"
        />
        <StatCard
          label="Customers"
          value={String(settlement.uniqueCustomers)}
          hint="Unique buyers via PayLater"
        />
      </div>
    </div>
  )
}

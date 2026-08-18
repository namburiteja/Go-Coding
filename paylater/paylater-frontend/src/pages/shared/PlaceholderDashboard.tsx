import { Card } from '../../components/ui/Card'

interface PlaceholderDashboardProps {
  roleLabel: string
  message: string
}

export function PlaceholderDashboard({ roleLabel, message }: PlaceholderDashboardProps) {
  return (
    <Card>
      <h2 className="text-xl font-semibold text-slate-900">{roleLabel} dashboard</h2>
      <p className="mt-2 text-sm text-slate-600">{message}</p>
      <p className="mt-4 text-xs text-slate-500">
        Full dashboard features will be added in later phases.
      </p>
    </Card>
  )
}

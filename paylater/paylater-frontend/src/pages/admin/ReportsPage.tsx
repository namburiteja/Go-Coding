import { ReportsPanel } from '../../components/admin/ReportsPanel'
import { PageHeader } from '../../components/ui/PageHeader'

export function AdminReportsPage() {
  return (
    <div>
      <PageHeader
        title="Reports"
        description="Credit-limit, outstanding due, customer lookup, and merchant fee reports."
      />
      <ReportsPanel />
    </div>
  )
}

import type { CustomerStatus } from '../../types/customer'

interface StatusBadgeProps {
  status: CustomerStatus | string
}

export function StatusBadge({ status }: StatusBadgeProps) {
  const normalized = status.toUpperCase()
  const isActive = normalized === 'ACTIVE'

  return (
    <span
      className={`inline-flex items-center rounded-full px-2.5 py-1 text-xs font-semibold ${
        isActive
          ? 'bg-emerald-100 text-emerald-800'
          : 'bg-red-100 text-red-800'
      }`}
    >
      <span
        className={`mr-1.5 h-1.5 w-1.5 rounded-full ${isActive ? 'bg-emerald-500' : 'bg-red-500'}`}
      />
      {normalized}
    </span>
  )
}

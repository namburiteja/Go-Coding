interface AlertProps {
  variant?: 'error' | 'success' | 'info'
  children: string
}

const styles = {
  error: 'border-red-200 bg-red-50 text-red-700',
  success: 'border-emerald-200 bg-emerald-50 text-emerald-700',
  info: 'border-sky-200 bg-sky-50 text-sky-700',
}

export function Alert({ variant = 'info', children }: AlertProps) {
  return (
    <div className={`rounded-lg border px-3 py-2 text-sm ${styles[variant]}`} role="alert">
      {children}
    </div>
  )
}

import { Link } from 'react-router-dom'
import type { ReactNode } from 'react'

interface PublicLayoutProps {
  children: ReactNode
  title?: string
  subtitle?: string
}

export function PublicLayout({ children, title, subtitle }: PublicLayoutProps) {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-100 via-slate-50 to-sky-50">
      <header className="border-b border-slate-200/80 bg-white/80 backdrop-blur">
        <div className="mx-auto flex max-w-5xl items-center justify-between px-4 py-4">
          <Link to="/" className="text-lg font-semibold tracking-tight text-slate-900">
            PayLater
          </Link>
          <nav className="flex gap-3 text-sm text-slate-600">
            <Link className="hover:text-slate-900" to="/customer/login">
              Customer
            </Link>
            <Link className="hover:text-slate-900" to="/merchant/login">
              Merchant
            </Link>
            <Link className="hover:text-slate-900" to="/admin/login">
              Admin
            </Link>
          </nav>
        </div>
      </header>

      <main className="mx-auto flex max-w-lg flex-col gap-6 px-4 py-10">
        {(title || subtitle) && (
          <div className="space-y-1 text-center">
            {title ? <h1 className="text-2xl font-semibold text-slate-900">{title}</h1> : null}
            {subtitle ? <p className="text-sm text-slate-600">{subtitle}</p> : null}
          </div>
        )}
        {children}
      </main>
    </div>
  )
}

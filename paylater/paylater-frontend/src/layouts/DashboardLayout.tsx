import { NavLink, Navigate, Outlet, useNavigate } from 'react-router-dom'
import { Button } from '../components/ui/Button'
import { useAuth } from '../context/useAuth'
import type { Role } from '../types/auth'
import { dashboardPathForRole } from '../utils/jwt'

interface DashboardLayoutProps {
  role: Role
  title: string
}

const navByRole: Record<Role, { to: string; label: string }[]> = {
  CUSTOMER: [
    { to: '/customer/dashboard', label: 'Dashboard' },
    { to: '/customer/purchase', label: 'Purchase' },
    { to: '/customer/payback', label: 'Payback' },
    { to: '/customer/transactions', label: 'Transactions' },
    { to: '/customer/profile', label: 'Profile' },
  ],
  MERCHANT: [
    { to: '/merchant/dashboard', label: 'Dashboard' },
    { to: '/merchant/transactions', label: 'Sales' },
    { to: '/merchant/customers', label: 'Customers' },
    { to: '/merchant/reports', label: 'Settlement' },
    { to: '/merchant/profile', label: 'Profile' },
  ],
  ADMIN: [
    { to: '/admin/dashboard', label: 'Dashboard' },
    { to: '/admin/customers', label: 'Customers' },
    { to: '/admin/merchants', label: 'Merchants' },
    { to: '/admin/admins', label: 'Admins' },
    { to: '/admin/transactions', label: 'Transactions' },
    { to: '/admin/reports', label: 'Reports' },
    { to: '/admin/profile', label: 'Profile' },
  ],
}

export function DashboardLayout({ role, title }: DashboardLayoutProps) {
  const { user, logout } = useAuth()
  const navigate = useNavigate()

  if (!user) {
    return <Navigate to="/" replace />
  }

  if (user.role !== role) {
    return <Navigate to={dashboardPathForRole(user.role)} replace />
  }

  function handleLogout() {
    logout()
    navigate('/', { replace: true })
  }

  return (
    <div className="min-h-screen bg-slate-100 lg:flex">
      <aside className="border-b border-slate-200 bg-slate-900 text-white lg:sticky lg:top-0 lg:flex lg:h-screen lg:w-64 lg:flex-col lg:border-b-0 lg:border-r">
        <div className="px-5 py-5">
          <p className="text-xs uppercase tracking-[0.2em] text-slate-400">PayLater</p>
          <h1 className="mt-1 text-lg font-semibold">{title}</h1>
          <p className="mt-1 text-xs text-slate-400">Signed in as {user.role}</p>
        </div>

        <nav className="flex gap-1 overflow-x-auto px-3 pb-4 lg:flex-1 lg:flex-col lg:overflow-y-auto">
          {navByRole[role].map((item) => (
            <NavLink
              key={item.to}
              to={item.to}
              className={({ isActive }) =>
                `whitespace-nowrap rounded-lg px-3 py-2 text-sm transition ${
                  isActive
                    ? 'bg-white/15 font-medium text-white'
                    : 'text-slate-200 hover:bg-slate-800 hover:text-white'
                }`
              }
            >
              {item.label}
            </NavLink>
          ))}
        </nav>

        <div className="hidden px-4 pb-5 lg:block">
          <Button variant="secondary" className="w-full" onClick={handleLogout}>
            Logout
          </Button>
        </div>
      </aside>

      <div className="flex min-w-0 flex-1 flex-col">
        <header className="sticky top-0 z-10 flex items-center justify-between border-b border-slate-200 bg-white/95 px-4 py-3 backdrop-blur lg:px-8">
          <div>
            <p className="text-sm font-medium text-slate-900">{title} portal</p>
            <p className="text-xs text-slate-500">User ID #{user.userId}</p>
          </div>
          <Button variant="ghost" size="sm" className="lg:hidden" onClick={handleLogout}>
            Logout
          </Button>
        </header>
        <main className="flex-1 px-4 py-6 lg:px-8">
          <Outlet />
        </main>
      </div>
    </div>
  )
}

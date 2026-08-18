import { Link } from 'react-router-dom'
import { Card } from '../../components/ui/Card'
import { PublicLayout } from '../../layouts/PublicLayout'

const portals = [
  {
    title: 'Customer',
    description: 'Register, shop with credit, and repay on the 5th.',
    loginTo: '/customer/login',
    registerTo: '/customer/register',
    registerLabel: 'Register',
  },
  {
    title: 'Merchant',
    description: 'Accept PayLater purchases and track settlements.',
    loginTo: '/merchant/login',
    registerTo: '/merchant/register',
    registerLabel: 'Register',
  },
  {
    title: 'Admin',
    description: 'Manage customers, merchants, and platform reports.',
    loginTo: '/admin/login',
    registerTo: null,
    registerLabel: null,
  },
] as const

export function LandingPage() {
  return (
    <PublicLayout
      title="Welcome to PayLater"
      subtitle="One app for customers, merchants, and administrators"
    >
      <div className="grid gap-4">
        {portals.map((portal) => (
          <Card key={portal.title}>
            <h2 className="text-lg font-semibold text-slate-900">{portal.title}</h2>
            <p className="mt-1 text-sm text-slate-600">{portal.description}</p>
            <div className="mt-4 flex flex-wrap gap-2">
              <Link
                to={portal.loginTo}
                className="rounded-lg bg-slate-900 px-3 py-2 text-sm font-medium text-white hover:bg-slate-800"
              >
                Login
              </Link>
              {portal.registerTo ? (
                <Link
                  to={portal.registerTo}
                  className="rounded-lg border border-slate-300 bg-white px-3 py-2 text-sm font-medium text-slate-800 hover:bg-slate-50"
                >
                  {portal.registerLabel}
                </Link>
              ) : (
                <span className="rounded-lg px-3 py-2 text-xs text-slate-500">
                  No public registration
                </span>
              )}
            </div>
          </Card>
        ))}
      </div>
    </PublicLayout>
  )
}

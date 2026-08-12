import type { ReactNode } from 'react'
import { Navigate, Route, Routes } from 'react-router-dom'
import { useAuth } from '../context/useAuth'
import { DashboardLayout } from '../layouts/DashboardLayout'
import { AdminLoginPage } from '../pages/auth/AdminLoginPage'
import { CustomerLoginPage, CustomerRegisterPage } from '../pages/auth/CustomerAuthPages'
import { MerchantLoginPage, MerchantRegisterPage } from '../pages/auth/MerchantAuthPages'
import { LandingPage } from '../pages/public/LandingPage'
import { AdminAdminsPage } from '../pages/admin/AdminsPage'
import { AdminCustomersPage } from '../pages/admin/CustomersPage'
import { AdminDashboardPage } from '../pages/admin/DashboardPage'
import { AdminMerchantsPage } from '../pages/admin/MerchantsPage'
import { AdminProfilePage } from '../pages/admin/ProfilePage'
import { AdminReportsPage } from '../pages/admin/ReportsPage'
import { AdminTransactionsPage } from '../pages/admin/TransactionsPage'
import { CustomerDashboardPage } from '../pages/customer/DashboardPage'
import { CustomerPaybackPage } from '../pages/customer/PaybackPage'
import { CustomerProfilePage } from '../pages/customer/ProfilePage'
import { CustomerPurchasePage } from '../pages/customer/PurchasePage'
import { CustomerTransactionsPage } from '../pages/customer/TransactionsPage'
import { MerchantCustomersPage } from '../pages/merchant/CustomersPage'
import { MerchantDashboardPage } from '../pages/merchant/DashboardPage'
import { MerchantProfilePage } from '../pages/merchant/ProfilePage'
import { MerchantSettlementPage } from '../pages/merchant/SettlementPage'
import { MerchantTransactionsPage } from '../pages/merchant/TransactionsPage'
import { dashboardPathForRole } from '../utils/jwt'
import { ProtectedRoute } from './ProtectedRoute'
import { RoleRoute } from './RoleRoute'

function GuestOnly({ children }: { children: ReactNode }) {
  const { isAuthenticated, role } = useAuth()
  if (isAuthenticated && role) {
    return <Navigate to={dashboardPathForRole(role)} replace />
  }
  return children
}

export function AppRoutes() {
  return (
    <Routes>
      <Route
        path="/"
        element={
          <GuestOnly>
            <LandingPage />
          </GuestOnly>
        }
      />

      <Route
        path="/customer/login"
        element={
          <GuestOnly>
            <CustomerLoginPage />
          </GuestOnly>
        }
      />
      <Route
        path="/customer/register"
        element={
          <GuestOnly>
            <CustomerRegisterPage />
          </GuestOnly>
        }
      />
      <Route
        path="/merchant/login"
        element={
          <GuestOnly>
            <MerchantLoginPage />
          </GuestOnly>
        }
      />
      <Route
        path="/merchant/register"
        element={
          <GuestOnly>
            <MerchantRegisterPage />
          </GuestOnly>
        }
      />
      <Route
        path="/admin/login"
        element={
          <GuestOnly>
            <AdminLoginPage />
          </GuestOnly>
        }
      />

      <Route element={<ProtectedRoute />}>
        <Route element={<RoleRoute allow="CUSTOMER" />}>
          <Route path="/customer" element={<DashboardLayout role="CUSTOMER" title="Customer" />}>
            <Route index element={<Navigate to="dashboard" replace />} />
            <Route path="dashboard" element={<CustomerDashboardPage />} />
            <Route path="purchase" element={<CustomerPurchasePage />} />
            <Route path="payback" element={<CustomerPaybackPage />} />
            <Route path="transactions" element={<CustomerTransactionsPage />} />
            <Route path="profile" element={<CustomerProfilePage />} />
          </Route>
        </Route>

        <Route element={<RoleRoute allow="MERCHANT" />}>
          <Route path="/merchant" element={<DashboardLayout role="MERCHANT" title="Merchant" />}>
            <Route index element={<Navigate to="dashboard" replace />} />
            <Route path="dashboard" element={<MerchantDashboardPage />} />
            <Route path="customers" element={<MerchantCustomersPage />} />
            <Route path="transactions" element={<MerchantTransactionsPage />} />
            <Route path="reports" element={<MerchantSettlementPage />} />
            <Route path="profile" element={<MerchantProfilePage />} />
          </Route>
        </Route>

        <Route element={<RoleRoute allow="ADMIN" />}>
          <Route path="/admin" element={<DashboardLayout role="ADMIN" title="Admin" />}>
            <Route index element={<Navigate to="dashboard" replace />} />
            <Route path="dashboard" element={<AdminDashboardPage />} />
            <Route path="customers" element={<AdminCustomersPage />} />
            <Route path="merchants" element={<AdminMerchantsPage />} />
            <Route path="admins" element={<AdminAdminsPage />} />
            <Route path="transactions" element={<AdminTransactionsPage />} />
            <Route path="reports" element={<AdminReportsPage />} />
            <Route path="profile" element={<AdminProfilePage />} />
          </Route>
        </Route>
      </Route>

      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  )
}

import { Navigate, Outlet } from 'react-router-dom'
import { useAuth } from '../context/useAuth'
import type { Role } from '../types/auth'
import { dashboardPathForRole } from '../utils/jwt'

interface RoleRouteProps {
  allow: Role
}

export function RoleRoute({ allow }: RoleRouteProps) {
  const { isAuthenticated, role } = useAuth()

  if (!isAuthenticated || !role) {
    return <Navigate to="/" replace />
  }

  if (role !== allow) {
    return <Navigate to={dashboardPathForRole(role)} replace />
  }

  return <Outlet />
}

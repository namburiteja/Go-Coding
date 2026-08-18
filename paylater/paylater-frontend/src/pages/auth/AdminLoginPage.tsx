import { type FormEvent, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { Card } from '../../components/ui/Card'
import { Input } from '../../components/ui/Input'
import { useAuth } from '../../context/useAuth'
import { useApiError } from '../../hooks/useApiError'
import { PublicLayout } from '../../layouts/PublicLayout'
import { loginAdmin } from '../../services/api/auth.api'
import { dashboardPathForRole } from '../../utils/jwt'

export function AdminLoginPage() {
  const navigate = useNavigate()
  const { login } = useAuth()
  const { error, capture, clear } = useApiError()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)

  async function onSubmit(event: FormEvent) {
    event.preventDefault()
    clear()
    setLoading(true)
    try {
      const { token } = await loginAdmin({ email, password })
      login(token)
      setPassword('')
      navigate(dashboardPathForRole('ADMIN'), { replace: true })
    } catch (err) {
      capture(err, 'Admin login failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <PublicLayout
      title="Admin login"
      subtitle="Administrator access only. New admins are created from the Admin dashboard."
    >
      <Card>
        <form className="space-y-4" onSubmit={onSubmit}>
          {error ? <Alert variant="error">{error}</Alert> : null}
          <Input
            label="Email"
            type="email"
            name="email"
            autoComplete="email"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <Input
            label="Password"
            type="password"
            name="password"
            autoComplete="current-password"
            required
            minLength={6}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button type="submit" className="w-full" loading={loading}>
            Sign in as admin
          </Button>
        </form>
      </Card>
    </PublicLayout>
  )
}

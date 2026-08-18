import { type FormEvent, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Alert } from '../../components/ui/Alert'
import { Button } from '../../components/ui/Button'
import { Card } from '../../components/ui/Card'
import { Input } from '../../components/ui/Input'
import { useAuth } from '../../context/useAuth'
import { useApiError } from '../../hooks/useApiError'
import { PublicLayout } from '../../layouts/PublicLayout'
import { loginMerchant, registerMerchant } from '../../services/api/auth.api'
import { dashboardPathForRole } from '../../utils/jwt'

export function MerchantLoginPage() {
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
      const { token } = await loginMerchant({ email, password })
      login(token)
      setPassword('')
      navigate(dashboardPathForRole('MERCHANT'), { replace: true })
    } catch (err) {
      capture(err, 'Merchant login failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <PublicLayout title="Merchant login" subtitle="Manage sales and settlements">
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
            Sign in
          </Button>
        </form>
        <p className="mt-4 text-center text-sm text-slate-600">
          New merchant?{' '}
          <Link className="font-medium text-slate-900 underline" to="/merchant/register">
            Register your store
          </Link>
        </p>
      </Card>
    </PublicLayout>
  )
}

export function MerchantRegisterPage() {
  const navigate = useNavigate()
  const { error, capture, clear } = useApiError()
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [phone, setPhone] = useState('')
  const [password, setPassword] = useState('')
  const [success, setSuccess] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  async function onSubmit(event: FormEvent) {
    event.preventDefault()
    clear()
    setSuccess(null)
    setLoading(true)
    try {
      const result = await registerMerchant({ name, email, phone, password })
      setPassword('')
      setSuccess(result.message || 'Merchant registered successfully')
      setTimeout(() => navigate('/merchant/login'), 800)
    } catch (err) {
      capture(err, 'Merchant registration failed')
    } finally {
      setLoading(false)
    }
  }

  return (
    <PublicLayout title="Merchant register" subtitle="Join the PayLater merchant network">
      <Card>
        <form className="space-y-4" onSubmit={onSubmit}>
          {error ? <Alert variant="error">{error}</Alert> : null}
          {success ? <Alert variant="success">{success}</Alert> : null}
          <Input
            label="Business name"
            name="name"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
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
            label="Phone"
            name="phone"
            required
            value={phone}
            onChange={(e) => setPhone(e.target.value)}
          />
          <Input
            label="Password"
            type="password"
            name="password"
            autoComplete="new-password"
            required
            minLength={6}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button type="submit" className="w-full" loading={loading}>
            Create merchant account
          </Button>
        </form>
        <p className="mt-4 text-center text-sm text-slate-600">
          Already registered?{' '}
          <Link className="font-medium text-slate-900 underline" to="/merchant/login">
            Sign in
          </Link>
        </p>
      </Card>
    </PublicLayout>
  )
}

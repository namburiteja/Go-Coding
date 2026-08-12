import { type FormEvent, useState } from 'react'
import { getErrorMessage } from '../../types/api'
import { registerAdmin } from '../../services/api/admin.api'
import { Alert } from '../ui/Alert'
import { Button } from '../ui/Button'
import { Card } from '../ui/Card'
import { Input } from '../ui/Input'

interface CreateAdminFormProps {
  onSuccess?: () => void
}

export function CreateAdminForm({ onSuccess }: CreateAdminFormProps) {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [success, setSuccess] = useState<string | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(false)

  async function onSubmit(event: FormEvent) {
    event.preventDefault()
    setError(null)
    setSuccess(null)

    if (!name.trim() || !email.trim() || !password) {
      setError('Name, email, and password are required.')
      return
    }
    if (password.length < 6) {
      setError('Password must be at least 6 characters.')
      return
    }

    setLoading(true)
    try {
      const result = await registerAdmin({
        name: name.trim(),
        email: email.trim(),
        password,
      })
      setSuccess(result.message || 'Admin registered successfully')
      setName('')
      setEmail('')
      setPassword('')
      onSuccess?.()
    } catch (err) {
      setError(getErrorMessage(err, 'Failed to register admin'))
    } finally {
      setLoading(false)
    }
  }

  return (
    <Card>
      <h3 className="text-base font-semibold text-slate-900">Create admin</h3>
      <p className="mt-1 text-sm text-slate-600">
        Registration is protected — only authenticated admins can create new admins. There is no
        public signup.
      </p>
      <form className="mt-4 space-y-4" onSubmit={onSubmit} autoComplete="off">
        {error ? <Alert variant="error">{error}</Alert> : null}
        {success ? <Alert variant="success">{success}</Alert> : null}
        <Input
          label="Name"
          name="admin-name"
          required
          value={name}
          onChange={(e) => setName(e.target.value)}
          autoComplete="off"
        />
        <Input
          label="Email"
          type="email"
          name="admin-email"
          required
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          autoComplete="off"
        />
        <Input
          label="Temporary password"
          type="password"
          name="admin-password"
          required
          minLength={6}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          autoComplete="new-password"
        />
        <Button type="submit" loading={loading}>
          Register admin
        </Button>
      </form>
    </Card>
  )
}

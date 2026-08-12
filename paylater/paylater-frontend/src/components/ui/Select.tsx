import type { SelectHTMLAttributes } from 'react'

interface SelectProps extends SelectHTMLAttributes<HTMLSelectElement> {
  label: string
  error?: string
  options: { value: string; label: string }[]
  placeholder?: string
}

export function Select({
  label,
  error,
  options,
  placeholder = 'Select…',
  id,
  className = '',
  ...props
}: SelectProps) {
  const selectId = id || props.name || label.toLowerCase().replace(/\s+/g, '-')

  return (
    <label className="block space-y-1.5" htmlFor={selectId}>
      <span className="text-sm font-medium text-slate-700">{label}</span>
      <select
        id={selectId}
        className={`w-full rounded-lg border border-slate-300 bg-white px-3 py-2.5 text-sm text-slate-900 shadow-sm outline-none transition focus:border-slate-500 focus:ring-2 focus:ring-slate-200 ${error ? 'border-red-400 focus:border-red-500 focus:ring-red-100' : ''} ${className}`}
        {...props}
      >
        <option value="">{placeholder}</option>
        {options.map((opt) => (
          <option key={opt.value} value={opt.value}>
            {opt.label}
          </option>
        ))}
      </select>
      {error ? <span className="text-xs text-red-600">{error}</span> : null}
    </label>
  )
}

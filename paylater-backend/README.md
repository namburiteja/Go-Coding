# PayLater

Go workspace monorepo for a pay-later backend: five domain services behind a pure API gateway.

## Layout

```
paylater/
├── services/
│   ├── admin/       # :9091 — auth, admin ops
│   ├── merchant/    # :9092 — merchants
│   ├── customer/    # :9093 — customers & credit
│   ├── ledger/      # :9094 — transactions & payments
│   └── report/      # :9095 — analytics / reports
├── gateway/         # :9090 — edge proxy only (no domain DB)
├── shared/
├── deployments/
└── README.md
```

## Architecture

- **Gateway** is a pure edge: routing and proxy to downstream services only.
- Domain logic and SQL live in the five services under `services/`.
- Each process loads **only its own** `.env` (under `gateway/` or `services/<name>/`).
- All services still use the shared MySQL database `paylater` (`DB_NAME`). Database-per-service comes in a later phase.
- Use the **same** `JWT_SECRET` in every service that issues or validates tokens.

## Configuration

| Process | Config file | Notes |
|---------|-------------|--------|
| Gateway | `gateway/.env` | `PORT` + upstream `*_SERVICE_URL` only (no DB) |
| Admin | `services/admin/.env` | `PORT`, `DB_*`, `JWT_*`, `INTERNAL_SERVICE_TOKEN` |
| Merchant | `services/merchant/.env` | same pattern |
| Customer | `services/customer/.env` | same pattern |
| Ledger | `services/ledger/.env` | plus `CUSTOMER_SERVICE_URL`, `MERCHANT_SERVICE_URL` |
| Report | `services/report/.env` | `PORT`, `DB_*`, `JWT_SECRET`, `INTERNAL_SERVICE_TOKEN` |

Copy each `.env.example` → `.env` before the first run (examples are committed; `.env` is gitignored).

## Run

From the **repo root** (each binary resolves `services/<name>/.env` or `gateway/.env`):

```bash
go run ./services/admin/cmd
go run ./services/merchant/cmd
go run ./services/customer/cmd
go run ./services/ledger/cmd
go run ./services/report/cmd
go run ./gateway/cmd
```

Or from a service directory:

```bash
cd services/admin && go run ./cmd
```

Gateway proxies admin, merchant, customer, ledger, and report.

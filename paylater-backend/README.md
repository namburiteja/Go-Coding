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

## Architecture (Phase 7)

- **Gateway** is a pure edge: routing and proxy to downstream services only.
- Domain logic and SQL live in the five services under `services/`.
- Shared MySQL is configured via root `.env` (`DB_*`); each service also has its own `.env.example`.

## Run

From the **repo root** (load `.env` as needed):

```bash
go run ./services/admin/cmd
go run ./services/merchant/cmd
go run ./services/customer/cmd
go run ./services/ledger/cmd
go run ./services/report/cmd
go run ./gateway/cmd
```

Gateway proxies admin, merchant, customer, ledger, and report.
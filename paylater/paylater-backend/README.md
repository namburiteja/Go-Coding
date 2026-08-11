# PayLater

Go workspace monorepo for a pay-later backend: five domain services behind a pure API gateway.

## Layout

```
paylater/
├── services/
│   ├── admin/       # :9091 — auth, admin ops → admin_db
│   ├── merchant/    # :9092 — merchants → merchant_db
│   ├── customer/    # :9093 — customers & credit → customer_db
│   ├── ledger/      # :9094 — transactions & payments → ledger_db
│   └── report/      # :9095 — analytics aggregator (no database)
├── gateway/         # :9090 — edge proxy only (no domain DB)
├── shared/
├── deployments/
└── README.md
```

## Architecture

- **Gateway** is a pure edge: routing and proxy to downstream services only.
- Domain logic and SQL live under `services/` (except Report).
- Each process loads **only its own** `.env` (under `gateway/` or `services/<name>/`).
- **Database-per-service:** Admin, Merchant, Customer, and Ledger each connect to a dedicated MySQL database. No service may query another service's database.
- **Report** is an aggregator: no MySQL connection; it calls Customer/Merchant/Ledger over authenticated internal HTTP.
- Cross-domain reads/writes use internal REST (`X-Internal-Service-Token`), never cross-database SQL.
- Use the **same** `JWT_SECRET` in every service that issues or validates tokens.
- Use the **same** `INTERNAL_SERVICE_TOKEN` for all service-to-service calls.

## Databases

| Service | `DB_NAME` | Schema file |
|---------|-----------|-------------|
| Admin | `admin_db` | `services/admin/sql/schema/admins.sql` |
| Merchant | `merchant_db` | `services/merchant/sql/schema/merchants.sql` |
| Customer | `customer_db` | `services/customer/sql/schema/customers.sql` |
| Ledger | `ledger_db` | `services/ledger/sql/schema/transactions.sql` |
| Report | — | none |
| Gateway | — | none |

With Docker Compose, MySQL databases and tables are created automatically from `deployments/mysql/init/` on a fresh volume. For non-Docker setups, see `deployments/README.md`.

## Configuration

| Process | Config file | Notes |
|---------|-------------|--------|
| Gateway | `gateway/.env` | `PORT` + upstream `*_SERVICE_URL` only (no DB) |
| Admin | `services/admin/.env` | `PORT`, `DB_*` (`DB_NAME=admin_db`), `JWT_*` |
| Merchant | `services/merchant/.env` | `DB_NAME=merchant_db`, `INTERNAL_SERVICE_TOKEN` |
| Customer | `services/customer/.env` | `DB_NAME=customer_db`, `INTERNAL_SERVICE_TOKEN` |
| Ledger | `services/ledger/.env` | `DB_NAME=ledger_db` + `CUSTOMER_SERVICE_URL`, `MERCHANT_SERVICE_URL` |
| Report | `services/report/.env` | `PORT`, `JWT_SECRET`, upstream URLs, `INTERNAL_SERVICE_TOKEN` (no `DB_*`) |

Copy each `.env.example` → `.env` before the first run (examples are committed; `.env` is gitignored). Defaults are Docker Compose–ready (`DB_HOST=mysql`, Compose service names). Commented localhost alternatives are in each example for `go run` on the host.

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

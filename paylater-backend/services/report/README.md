# Report Service

Admin analytics APIs over shared MySQL read access to customer/merchant/transaction tables.

## Configuration

Copy `.env.example` → `.env` in this directory.

| Variable | Purpose |
|----------|---------|
| `PORT` | Listen port |
| `DB_*` | MySQL connection (shared `paylater` DB for now) |
| `JWT_SECRET` | Must match other services (admin-only routes) |
| `INTERNAL_SERVICE_TOKEN` | Reserved for Phase 10; unused today |

## Run

From this directory:

```bash
cp .env.example .env   # first time only
go run ./cmd
```

From repo root:

```bash
go run ./services/report/cmd
```

## Public paths (proxied by gateway)

- `GET /reports/credit-limit`
- `GET /reports/customers-due`
- `GET /reports/customer-due/:name`
- `GET /reports/merchant-fees`

## Docker

```bash
docker build -f services/report/Dockerfile .
```

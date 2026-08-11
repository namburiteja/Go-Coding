# Report Service

Admin analytics aggregator. Does **not** own a database — it calls Customer, Merchant, and Ledger over authenticated internal HTTP and aggregates in Go.

## Configuration

Copy `.env.example` → `.env` in this directory.

| Variable | Purpose |
|----------|---------|
| `PORT` | Listen port |
| `JWT_SECRET` | Must match other services (admin-only public routes) |
| `CUSTOMER_SERVICE_URL` | Upstream for customer report data |
| `MERCHANT_SERVICE_URL` | Upstream for merchant names |
| `LEDGER_SERVICE_URL` | Upstream for merchant fee totals |
| `INTERNAL_SERVICE_TOKEN` | Attached by `shared/httpclient` on S2S calls |

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

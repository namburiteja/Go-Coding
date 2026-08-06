# Ledger Service

Owns the `transactions` table and purchase/payback/transaction-list APIs.

## Configuration

Copy `.env.example` → `.env` in this directory.

| Variable | Purpose |
|----------|---------|
| `PORT` | Listen port |
| `DB_*` | MySQL connection (shared `paylater` DB for now) |
| `JWT_SECRET` | Must match other services (validates customer/merchant/admin JWTs) |
| `CUSTOMER_SERVICE_URL` | Base URL for credit/due/block internal APIs |
| `MERCHANT_SERVICE_URL` | Base URL for commission internal API |
| `INTERNAL_SERVICE_TOKEN` | Attached automatically by `shared/httpclient` on S2S calls |

## Run

From this directory:

```bash
cp .env.example .env   # first time only
go run ./cmd
```

From repo root:

```bash
go run ./services/ledger/cmd
```

## Public paths (proxied by gateway)

- `POST /customers/purchase`
- `POST /customers/payback`
- `GET /customers/me/transactions`
- `GET /merchants/me/transactions`
- `GET /transactions`

## Internal API (Report)

Require header `X-Internal-Service-Token: <INTERNAL_SERVICE_TOKEN>`.

- `GET /internal/transactions/reports/merchant-fees` — `{merchant_id, total_fee_collected}` aggregates

## Docker

```bash
docker build -f services/ledger/Dockerfile .
```

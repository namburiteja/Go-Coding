# Merchant Service

Independent microservice owning the `merchants` table and all `/merchants/*` public APIs.

## Run

From repo root (so shared `.env` / JWT secret loads):

```bash
go run ./services/merchant/cmd
```

Or from this directory after copying `.env.example` → `.env`:

```bash
go run ./cmd
```

Default address: `:9092` (`MERCHANT_SERVICE_ADDR`).

## Internal API (Ledger)

`GET /internal/merchants/:id/commission` — commission snapshot for purchase fees (not proxied by gateway).

## SQLC

```bash
sqlc generate
```

## Docker

From repo root:

```bash
docker build -f services/merchant/Dockerfile .
```

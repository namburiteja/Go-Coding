# Merchant Service

Independent microservice owning the `merchants` table and all `/merchants/*` public APIs.

## Configuration

Copy `.env.example` → `.env` in this directory.

| Variable | Purpose |
|----------|---------|
| `PORT` | Listen port |
| `DB_*` | MySQL connection (shared `paylater` DB for now) |
| `JWT_SECRET` | Must match other services |
| `JWT_EXPIRY` | Token lifetime for merchant login |
| `INTERNAL_SERVICE_TOKEN` | Reserved for Phase 10; unused today |

## Run

From this directory:

```bash
cp .env.example .env   # first time only
go run ./cmd
```

From repo root:

```bash
go run ./services/merchant/cmd
```

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

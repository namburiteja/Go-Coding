# Customer Service

Independent microservice owning the `customers` table and `/customers/*` public APIs.

## Configuration

Copy `.env.example` → `.env` in this directory.

| Variable | Purpose |
|----------|---------|
| `PORT` | Listen port |
| `DB_*` | MySQL connection (shared `paylater` DB for now) |
| `JWT_SECRET` | Must match other services |
| `JWT_EXPIRY` | Token lifetime for customer login |
| `INTERNAL_SERVICE_TOKEN` | Reserved for Phase 10; unused today |

## Run

From this directory:

```bash
cp .env.example .env   # first time only
go run ./cmd
```

From repo root:

```bash
go run ./services/customer/cmd
```

## Internal APIs (Ledger)

- `GET /internal/customers/:id/credit` — credit snapshot (`FOR UPDATE` read)
- `PUT /internal/customers/:id/due` — body `{"total_due":"12.34"}`
- `PUT /internal/customers/:id/block` — set status `BLOCKED`

## Docker

From repo root:

```bash
docker build -f services/customer/Dockerfile .
```

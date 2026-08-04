# Customer Service

Independent microservice owning the `customers` table and `/customers/*` public APIs.

## Run

From repo root:

```bash
go run ./services/customer/cmd
```

Default address: `:9093` (`CUSTOMER_SERVICE_ADDR`).

## Internal APIs (Ledger)

- `GET /internal/customers/:id/credit` — credit snapshot (`FOR UPDATE` read)
- `PUT /internal/customers/:id/due` — body `{"total_due":"12.34"}`
- `PUT /internal/customers/:id/block` — set status `BLOCKED`

## Docker

From repo root:

```bash
docker build -f services/customer/Dockerfile .
```

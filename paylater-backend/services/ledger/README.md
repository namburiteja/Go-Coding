# Ledger Service

Owns the `transactions` table and purchase/payback/transaction-list APIs.

## Run

From repo root:

```bash
go run ./services/ledger/cmd
```

Default address: `:9094` (`LEDGER_SERVICE_ADDR`).

Requires Customer (`CUSTOMER_SERVICE_URL`) and Merchant (`MERCHANT_SERVICE_URL`) for credit/commission HTTP calls.

## Public paths (proxied by gateway)

- `POST /customers/purchase`
- `POST /customers/payback`
- `GET /customers/me/transactions`
- `GET /merchants/me/transactions`
- `GET /transactions`

## Docker

```bash
docker build -f services/ledger/Dockerfile .
```

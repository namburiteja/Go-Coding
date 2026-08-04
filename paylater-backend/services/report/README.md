# Report Service

Admin analytics APIs over shared MySQL read access to customer/merchant/transaction tables.

## Run

From repo root:

```bash
go run ./services/report/cmd
```

Default address: `:9095` (`REPORT_SERVICE_ADDR`).

## Public paths (proxied by gateway)

- `GET /reports/credit-limit`
- `GET /reports/customers-due`
- `GET /reports/customer-due/:name`
- `GET /reports/merchant-fees`

## Docker

```bash
docker build -f services/report/Dockerfile .
```

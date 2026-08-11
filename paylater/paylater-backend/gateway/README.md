# Gateway

Pure API edge. Proxies all public routes to domain services — no domain DB or JWT validation.

## Configuration

Copy `.env.example` → `.env` in this directory.

| Variable | Purpose |
|----------|---------|
| `PORT` | Listen port (default example `9090`) |
| `ADMIN_SERVICE_URL` | Upstream admin service |
| `MERCHANT_SERVICE_URL` | Upstream merchant service |
| `CUSTOMER_SERVICE_URL` | Upstream customer service |
| `LEDGER_SERVICE_URL` | Upstream ledger service |
| `REPORT_SERVICE_URL` | Upstream report service |

No `DB_*` or `JWT_*` — the gateway does not touch the database or validate tokens.

## Run

From repo root:

```bash
go run ./gateway/cmd
```

From this directory:

```bash
go run ./cmd
```

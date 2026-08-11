# Admin service

Standalone admin microservice (JWT admin auth, admin CRUD).

## Configuration

Copy `.env.example` → `.env` in this directory.

| Variable | Purpose |
|----------|---------|
| `PORT` | Listen port |
| `DB_*` | MySQL connection (`DB_NAME=admin_db`, owns `admins` only) |
| `JWT_SECRET` | Must match other services |
| `JWT_EXPIRY` | Token lifetime (e.g. `24h`) |
| `INTERNAL_SERVICE_TOKEN` | Shared S2S token (for future internal APIs) |

## Run locally

From this directory:

```bash
cp .env.example .env   # first time only
go run ./cmd
```

From repo root:

```bash
go run ./services/admin/cmd
```

## Docker

Build from **repo root** (needs `shared/` for the replace directive):

```bash
docker build -f services/admin/Dockerfile .
```

## sqlc

```bash
sqlc generate
```

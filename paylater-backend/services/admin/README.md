# Admin service

Standalone admin microservice (JWT admin auth, admin CRUD).

## Run locally

```bash
cd services/admin
cp .env.example .env
go run ./cmd
```

Default listen address: `:9091`.

## Docker

Build from **repo root** (needs `shared/` for the replace directive):

```bash
docker build -f services/admin/Dockerfile .
```

## sqlc

```bash
sqlc generate
```

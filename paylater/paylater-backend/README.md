# PayLater

Go workspace monorepo for a pay-later backend: five domain services behind a pure API gateway, plus a React SPA.

## Layout

```
paylater/
├── paylater-backend/
│   ├── services/
│   │   ├── admin/       # :9091 — auth, admin ops → admin_db
│   │   ├── merchant/    # :9092 — merchants → merchant_db
│   │   ├── customer/    # :9093 — customers & credit → customer_db
│   │   ├── ledger/      # :9094 — transactions & payments → ledger_db
│   │   └── report/      # :9095 — analytics aggregator (no database)
│   ├── gateway/         # :9090 — edge proxy only (no domain DB)
│   ├── shared/
│   ├── deployments/
│   └── docker-compose.yml
└── paylater-frontend/   # Docker host :8081 (Jenkins keeps :8080)
```

## Architecture

- **Gateway** is a pure edge: routing, CORS, and proxy to downstream services only.
- Domain logic and SQL live under `services/` (except Report).
- Each process loads **only its own** `.env` (under `gateway/` or `services/<name>/`), or uses process env (Docker/Kubernetes).
- **Database-per-service:** Admin, Merchant, Customer, and Ledger each connect to a dedicated MySQL database.
- **Report** is an aggregator: no MySQL; it calls Customer/Merchant/Ledger over authenticated internal HTTP.
- Cross-domain reads/writes use internal REST (`X-Internal-Service-Token`), never cross-database SQL.
- Use the **same** `JWT_SECRET` in every service that issues or validates tokens.
- Use the **same** `INTERNAL_SERVICE_TOKEN` for all service-to-service calls.
- Inside Docker/K8s, use **service names** (e.g. `http://customer:9093`), not `localhost` or container IPs.

## Ports (host)

| Service | Host port |
|---------|-----------|
| Jenkins | **8080** (reserved — do not bind PayLater here) |
| Frontend (Docker) | **8081** |
| Gateway | **9090** |
| Admin | **9091** |
| Merchant | **9092** |
| Customer | **9093** |
| Ledger | **9094** |
| Report | **9095** |
| MySQL | **3307** → container 3306 |

## Databases

| Service | `DB_NAME` | Schema file |
|---------|-----------|-------------|
| Admin | `admin_db` | `services/admin/sql/schema/admins.sql` |
| Merchant | `merchant_db` | `services/merchant/sql/schema/merchants.sql` |
| Customer | `customer_db` | `services/customer/sql/schema/customers.sql` |
| Ledger | `ledger_db` | `services/ledger/sql/schema/transactions.sql` |
| Report | — | none |
| Gateway | — | none |

With Docker Compose, MySQL databases and tables are created from `deployments/mysql/init/` on a fresh volume. See `deployments/README.md`.

## Configuration

| Process | Config file | Notes |
|---------|-------------|--------|
| Gateway | `gateway/.env` | `PORT`, upstream `*_SERVICE_URL`, `CORS_ALLOWED_ORIGINS` |
| Admin | `services/admin/.env` | `PORT`, `DB_*`, `JWT_*` |
| Merchant | `services/merchant/.env` | `DB_NAME=merchant_db`, `INTERNAL_SERVICE_TOKEN` |
| Customer | `services/customer/.env` | `DB_NAME=customer_db`, `INTERNAL_SERVICE_TOKEN` |
| Ledger | `services/ledger/.env` | `DB_*` + `CUSTOMER_SERVICE_URL`, `MERCHANT_SERVICE_URL` |
| Report | `services/report/.env` | upstream URLs, `INTERNAL_SERVICE_TOKEN` (no `DB_*`) |
| Frontend | `paylater-frontend/.env` | `VITE_API_BASE_URL` (local Vite only) |

Copy each `.env.example` → `.env` before the first run. Defaults are Docker Compose–ready (`DB_HOST=mysql`, Compose service names). Commented localhost alternatives are in each example for `go run` on the host.

OS/Docker/Kubernetes environment variables are never overridden by `.env` files (`godotenv` load order), so ConfigMaps/Secrets work as expected.

## Admin bootstrap

Admin registration (`POST /admins/register`) requires an existing **ADMIN** JWT (chicken-and-egg).

On a **fresh** MySQL volume, `deployments/mysql/init/06-admin-bootstrap.sql` seeds:

- Email: `admin@paylater.local`
- Password: `Admin@123`

Change this password after first login. To regenerate a bcrypt hash:

```bash
cd paylater-backend
docker run --rm -v "$(pwd)/scripts:/work" -w /work golang:1.26-alpine \
  sh -c 'go mod init genhash >/dev/null 2>&1; go get golang.org/x/crypto/bcrypt >/dev/null 2>&1; go run genhash.go YourNewPassword'
```

If the volume already exists without an admin, insert manually into `admin_db.admins` (see `deployments/README.md`).

## Docker Compose (recommended)

From `paylater-backend/`:

```bash
# Ensure each service .env exists (copy from .env.example if needed)
docker compose up --build -d
```

- UI: http://localhost:8081  
- Gateway: http://localhost:9090  

Frontend nginx proxies `/api/*` to `gateway:9090`. Images build locally (`paylater-*:local`).

## Local `go run`

```bash
go run ./services/admin/cmd
go run ./services/merchant/cmd
go run ./services/customer/cmd
go run ./services/ledger/cmd
go run ./services/report/cmd
go run ./gateway/cmd
```

Point `DB_HOST=localhost` / `DB_PORT=3307` and `*_SERVICE_URL=http://localhost:...` as commented in each `.env.example`.

## Tests & build

The workspace root is not a Go module, so `go test ./...` from `paylater-backend/` fails. Use:

```bash
./scripts/test-all.sh
# or: (cd gateway && go test ./...)  # repeat per module
```

```bash
go build -o /tmp/gateway ./gateway/cmd
# similarly for services/*/cmd
```

## Kubernetes readiness (no manifests yet)

- Env-based config; no hardcoded container IPs
- Service-to-service URLs use DNS names
- Independently buildable Docker images
- Stateless app containers; MySQL state in a volume/PVC
- Secrets (`JWT_SECRET`, `INTERNAL_SERVICE_TOKEN`, `DB_PASSWORD`) live in env / future Secrets, not source

Do not create Kubernetes YAML until explicitly requested.

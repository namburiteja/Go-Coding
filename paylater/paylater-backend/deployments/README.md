# Deployments

Docker Compose and MySQL initialization for PayLater.

## Database-per-service

| Database | Owned by | Tables |
|----------|----------|--------|
| `admin_db` | Admin service | `admins` |
| `merchant_db` | Merchant service | `merchants` |
| `customer_db` | Customer service | `customers` |
| `ledger_db` | Ledger service | `transactions` |

Report and Gateway have no database. Cross-domain data access is HTTP-only.

## Docker Compose (recommended)

From `paylater-backend/`:

1. Copy each `.env.example` → `.env` (see root `.env.example` for the list).
2. `docker compose up --build -d`

Host ports: gateway **9090**, frontend **8081**, MySQL **3307**. Do **not** bind PayLater to **8080** (Jenkins).

On a **fresh** MySQL volume, scripts in `mysql/init/` run once and create:

- the four databases + `go_user` (`01-init.sql`)
- tables in each database (`02`–`05-*-schema.sql`)
- bootstrap admin (`06-admin-bootstrap.sql`) — `admin@paylater.local` / `Admin@123`

Compose manages the named volume `mysql-data`. Init scripts run only when the data directory is empty.

### Frontend

The `frontend` service builds `../paylater-frontend` with `VITE_API_BASE_URL=/api` and serves it via nginx. Browser → `http://localhost:8081/api/...` → `gateway:9090/...`.

### Recovery if databases exist but tables are missing

This happens when an older broken init left a volume half-initialized. Init will **not** re-run.

Only if you can afford to lose that MySQL data (or have a backup):

```bash
docker compose down
docker volume rm paylater-backend_mysql-data
# volume name may vary; check with: docker volume ls | grep mysql
docker compose up --build -d
```

Do **not** delete the volume if it holds data you need.

### Manual admin bootstrap (existing volume)

```bash
docker exec -i mysql mysql -uroot -proot123 admin_db <<'SQL'
INSERT INTO admins (name, email, password)
SELECT 'Bootstrap Admin', 'admin@paylater.local',
  '$2a$10$N8gw78TqcLgcFSUbuuK0ze2ltSilG7gjljPrRezdTQ.F.RHtLeAtW'
WHERE NOT EXISTS (SELECT 1 FROM admins WHERE email = 'admin@paylater.local');
SQL
```

## Non-Docker / manual MySQL

If you run MySQL yourself (not via Compose init):

```sql
CREATE DATABASE admin_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE merchant_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE customer_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE ledger_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

```bash
mysql -u go_user -p admin_db    < ../services/admin/sql/schema/admins.sql
mysql -u go_user -p merchant_db < ../services/merchant/sql/schema/merchants.sql
mysql -u go_user -p customer_db < ../services/customer/sql/schema/customers.sql
mysql -u go_user -p ledger_db   < ../services/ledger/sql/schema/transactions.sql
```

Then insert a bootstrap admin (see above) or use `scripts/genhash.go` to hash a password.

## Kubernetes readiness

Compose uses env files and service DNS names so the same images can later run under Kubernetes with ConfigMaps/Secrets. No K8s manifests are checked in yet.

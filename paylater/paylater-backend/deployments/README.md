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
2. `docker compose up -d`

On a **fresh** MySQL volume, scripts in `mysql/init/` run once and create:

- the four databases + `go_user` (`01-init.sql`)
- tables in each database (`02`–`05-*-schema.sql`)

Compose manages the named volume `mysql-data`. Init scripts run only when the data directory is empty.

### Recovery if databases exist but tables are missing

This happens when an older broken init left a volume half-initialized. Init will **not** re-run.

Only if you can afford to lose that MySQL data (or have a backup):

```bash
docker compose down
docker volume rm paylater-backend_mysql-data
# volume name may vary; check with: docker volume ls | grep mysql
docker compose up -d
```

Do **not** delete the volume if it holds data you need.

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

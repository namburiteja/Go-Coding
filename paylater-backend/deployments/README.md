# Deployments

Placeholder for compose / Kubernetes manifests (later phases).

## Database-per-service (Phase 11)

MySQL databases are **not** created by the application. Create and initialize them manually before starting Admin, Merchant, Customer, or Ledger.

### 1. Create databases

```sql
CREATE DATABASE admin_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE merchant_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE customer_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE ledger_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Grant your app user access to each database (adjust user/host as needed):

```sql
GRANT ALL PRIVILEGES ON admin_db.* TO 'go_user'@'%';
GRANT ALL PRIVILEGES ON merchant_db.* TO 'go_user'@'%';
GRANT ALL PRIVILEGES ON customer_db.* TO 'go_user'@'%';
GRANT ALL PRIVILEGES ON ledger_db.* TO 'go_user'@'%';
FLUSH PRIVILEGES;
```

### 2. Apply schemas (from repository root)

```bash
mysql -u go_user -p admin_db    < services/admin/sql/schema/admins.sql
mysql -u go_user -p merchant_db < services/merchant/sql/schema/merchants.sql
mysql -u go_user -p customer_db < services/customer/sql/schema/customers.sql
mysql -u go_user -p ledger_db   < services/ledger/sql/schema/transactions.sql
```

### 3. Optional: migrate data from legacy shared `paylater` DB

If you previously used a single `paylater` database, see the migration steps in the Phase 11 completion notes (copy tables into the new databases; do not drop `paylater` until verified).

### Ownership

| Database | Owned by | Tables |
|----------|----------|--------|
| `admin_db` | Admin service | `admins` |
| `merchant_db` | Merchant service | `merchants` |
| `customer_db` | Customer service | `customers` |
| `ledger_db` | Ledger service | `transactions` |

Report and Gateway have no database. Cross-domain data access is HTTP-only.

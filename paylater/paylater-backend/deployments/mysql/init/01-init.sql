CREATE DATABASE IF NOT EXISTS admin_db;
CREATE DATABASE IF NOT EXISTS merchant_db;
CREATE DATABASE IF NOT EXISTS customer_db;
CREATE DATABASE IF NOT EXISTS ledger_db;

CREATE USER IF NOT EXISTS 'go_user'@'%' IDENTIFIED BY 'go123';

GRANT ALL PRIVILEGES ON admin_db.* TO 'go_user'@'%';
GRANT ALL PRIVILEGES ON merchant_db.* TO 'go_user'@'%';
GRANT ALL PRIVILEGES ON customer_db.* TO 'go_user'@'%';
GRANT ALL PRIVILEGES ON ledger_db.* TO 'go_user'@'%';

FLUSH PRIVILEGES;
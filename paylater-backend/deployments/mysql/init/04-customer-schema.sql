-- Owned by Customer service → database: customer_db
CREATE TABLE customers (
    id INT AUTO_INCREMENT PRIMARY KEY,

    name VARCHAR(50) NOT NULL,

    email VARCHAR(50) NOT NULL UNIQUE,

    password VARCHAR(255) NOT NULL,

    credit_limit DECIMAL(10,2) NOT NULL DEFAULT 2000,

    total_due DECIMAL(10,2) DEFAULT 0,

    payment_due_date DATE NOT NULL,

    status ENUM(
        'ACTIVE',
        'BLOCKED'
    ) DEFAULT 'ACTIVE',

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
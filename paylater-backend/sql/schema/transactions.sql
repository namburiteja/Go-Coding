CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,

    customer_id INT NOT NULL,
    merchant_id INT NULL,

    transaction_type ENUM('PURCHASE','PAYBACK') NOT NULL,

    amount DECIMAL(10,2) NOT NULL,

    commission_percentage DECIMAL(5,2),

    commission_amount DECIMAL(10,2),

    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);
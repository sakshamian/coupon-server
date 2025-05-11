TABLE:

Create table coupon(
	id int unsigned AUTO_INCREMENT PRIMARY KEY,
    coupon_code varchar(50) NOT NULL UNIQUE,
	usage_type ENUM('one_time', 'multi_use', 'time_based') NOT NULL,
	applicable_medicine_ids JSON,
    applicable_categories JSON,
    min_order_value DECIMAL(10, 2) NOT NULL, 
	valid_from DATETIME,
    valid_to DATETIME,
    terms_and_conditions TEXT,
    discount_type ENUM('inventory', 'charges') NOT NULL,
    discount_value DECIMAL(10, 2) NOT NULL,
    max_usage_per_user INT,
    is_active INT DEFAULT 1,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL
);
CREATE TABLE sensor_data (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sensor_value DOUBLE NOT NULL,
    sensor_type VARCHAR(50),
    id1 CHAR(3),
    id2 INT,
    timestamp DATETIME(6),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_id1_id2 (id1, id2), -- index from composite key
    INDEX idx_timestamp (timestamp) -- index on timestamp for faster queries
);
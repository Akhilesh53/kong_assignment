CREATE TABLE versions (
    id SERIAL PRIMARY KEY,
    service_id INT NOT NULL,
    version VARCHAR(50) NOT NULL,
    release_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE
);
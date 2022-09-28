CREATE TABLE IF NOT EXISTS affiliates(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    producer_id INT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES producers(id)
);
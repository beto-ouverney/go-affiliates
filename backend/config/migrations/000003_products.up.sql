CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    producer_id INT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES producers(id),
    UNIQUE (name, producer_id)
);
CREATE TABLE IF NOT EXISTS sales_producers(
    id SERIAL PRIMARY KEY,
    producer_id INT NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (producer_id) REFERENCES producers(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    value INT NOT NULL,
    commission INT,
    date TIMESTAMPTZ,
    UNIQUE (producer_id, product_id, date)
);
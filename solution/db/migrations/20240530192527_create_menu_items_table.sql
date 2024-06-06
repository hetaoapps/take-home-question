-- migrate:up
CREATE TABLE IF NOT EXISTS MenuItems (
    id VARCHAR(255) PRIMARY KEY,
    store_id VARCHAR(255) REFERENCES Stores(id),
    name VARCHAR(255) NOT NULL,
    price VARCHAR(255) NOT NULL
);


-- migrate:down
DROP TABLE MenuItems;

-- migrate:up
CREATE TABLE IF NOT EXISTS Recommendations (
    id SERIAL PRIMARY KEY,
    query TEXT NOT NULL,
    store_id VARCHAR(255) REFERENCES Stores(id)
);

-- migrate:down
DROP TABLE Recommendations;

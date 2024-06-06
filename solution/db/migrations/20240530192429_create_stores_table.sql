-- migrate:up
CREATE TABLE IF NOT EXISTS Stores IF NOT EXIST (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);


-- migrate:down
DROP TABLE Stores;
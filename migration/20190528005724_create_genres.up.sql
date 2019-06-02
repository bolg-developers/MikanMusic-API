CREATE TABLE IF NOT EXISTS genres (
        id CHAR(36) NOT NULL,
        name CHAR(255) NOT NULL UNIQUE,
        PRIMARY KEY(id)
);
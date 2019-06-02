CREATE TABLE IF NOT EXISTS musics (
        id CHAR(36) NOT NULL,
        PRIMARY KEY(id),
        artist_id CHAR(36),
        INDEX artist_id (artist_id),
        FOREIGN KEY (artist_id)
                REFERENCES users(id)
                ON DELETE SET NULL
                ON UPDATE RESTRICT,
        lyrist_id CHAR(36),
        INDEX lyrist_id (lyrist_id),
        FOREIGN KEY (lyrist_id)
                REFERENCES users(id)
                ON DELETE SET NULL
                ON UPDATE RESTRICT,
        composer_id CHAR(36),
        INDEX composer_id (composer_id),
        FOREIGN KEY (composer_id)
                REFERENCES users(id)
                ON DELETE SET NULL
                ON UPDATE RESTRICT,
        arranger_id CHAR(36),
        INDEX arranger_id (arranger_id),
        FOREIGN KEY (arranger_id)
                REFERENCES users(id)
                ON DELETE SET NULL
                ON UPDATE RESTRICT,
        name CHAR(255) NOT NULL,
        INDEX name (name),
        lyric VARCHAR(2000) NOT NULL,
        description VARCHAR(500) NOT NULL,
        count_like INT UNSIGNED NOT NULL,
        INDEX count_like (count_like),
        count_listen BIGINT UNSIGNED NOT NULL,
        INDEX count_listen (count_listen),
        resource_url CHAR(255) NOT NULL,
        artwork_url CHAR(255) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
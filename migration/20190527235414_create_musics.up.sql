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

INSERT INTO musics VALUES('c8b5749b-5486-4f7a-805e-1552dfb0284e', 'f2526355-03ba-472b-ab29-2688951d039f', 'f2526355-03ba-472b-ab29-2688951d039f', 'f2526355-03ba-472b-ab29-2688951d039f', 'f2526355-03ba-472b-ab29-2688951d039f', 'OMMC', 'いっきっきの～き～', '神曲です。', 0, 0, 'https://s3-ap-northeast-1.amazonaws.com/mikanmusic/3dcc4b84-84cf-42ba-b012-c10810894df4.mp3', 'https://s3-ap-northeast-1.amazonaws.com/mikanmusic/e445ec91-654b-41a2-b61b-00c0fbf5adf6.jpg', '2019-05-31 00:19:17', '2019-05-31 00:19:17');
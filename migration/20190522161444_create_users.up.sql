CREATE TABLE IF NOT EXISTS users (
        id CHAR(36) NOT NULL PRIMARY KEY,
        display_name CHAR(32) NOT NULL UNIQUE,
        icon_url CHAR(255) NOT NULL,
        introduction CHAR(255) DEFAULT '',
        email CHAR(255) NOT NULL UNIQUE,
        password CHAR(60) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users (id, display_name, icon_url, introduction, email, password, created_at, updated_at)
        VALUES('f2526355-03ba-472b-ab29-2688951d039f', '本田翼', 'https://s3-ap-northeast-1.amazonaws.com/mikanmusic/01f9e044-a3bd-46b5-a586-3ccfdfd01aad.jpg', 'ゲームすきです。', 'ilovegame@gmail.com', '$2a$10$I54htoPl2vx7Q1q0RLubAOffUbamgK3mvzKAPS7MbhsF41YjNGV6i', '2019-05-31 00:19:17', '2019-05-31 00:19:17');
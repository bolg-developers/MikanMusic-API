CREATE TABLE IF NOT EXISTS mikan_music.users (
        id CHAR(36) NOT NULL PRIMARY KEY,
        display_name CHAR(32) NOT NULL UNIQUE,
        icon_url CHAR(255) NOT NULL,
        introduction CHAR(255) DEFAULT '',
        email CHAR(255) NOT NULL UNIQUE,
        password CHAR(60) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        logout_at DATETIME
);
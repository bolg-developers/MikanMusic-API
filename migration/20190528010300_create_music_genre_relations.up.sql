CREATE TABLE IF NOT EXISTS music_genre_relations (
        music_id CHAR(36) NOT NULL,
        genre_id CHAR(36) NOT NULL,
        INDEX music_id (music_id),
        INDEX genre_id (genre_id),
        FOREIGN KEY(music_id)
                REFERENCES musics(id)
                ON DELETE CASCADE
                ON UPDATE RESTRICT,
        FOREIGN KEY(music_id)
                REFERENCES genres(id)
                ON DELETE CASCADE
                ON UPDATE RESTRICT
);
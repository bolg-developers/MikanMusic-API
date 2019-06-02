CREATE TABLE IF NOT EXISTS music_genre_relations (
        music_id CHAR(36) NOT NULL,
        genre_id CHAR(36) NOT NULL,
        INDEX music_id (music_id),
        INDEX genre_id (genre_id),
        FOREIGN KEY(music_id)
                REFERENCES musics(id)
                ON DELETE CASCADE
                ON UPDATE RESTRICT,
        FOREIGN KEY(genre_id)
                REFERENCES genres(id)
                ON DELETE CASCADE
                ON UPDATE RESTRICT
);

INSERT INTO music_genre_relations VALUES('c8b5749b-5486-4f7a-805e-1552dfb0284e', 'a1fb538b-f081-4362-80a9-9bc8cec053db');
INSERT INTO music_genre_relations VALUES('c8b5749b-5486-4f7a-805e-1552dfb0284e', '15c7d984-1065-4481-b7c0-228e833f9126');
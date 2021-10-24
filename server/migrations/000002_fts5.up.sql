CREATE VIRTUAL TABLE artist_fts USING fts5 (
    name, 
    content='artist',
    content_rowid='id',
    tokenize="trigram"
);

CREATE VIRTUAL TABLE album_fts USING fts5 (
    name, 
    content='album',
    content_rowid='id',
    tokenize="trigram"
);

CREATE VIRTUAL TABLE song_fts USING fts5 (
    name,
    lyrics,
    year,
    content='song',
    content_rowid='id',
    tokenize="trigram"
);

CREATE TRIGGER artist_ai AFTER INSERT ON artist
    BEGIN
        INSERT INTO artist_fts (rowid, name)
        VALUES (new.id, new.name);
    END;

CREATE TRIGGER artist_ad AFTER DELETE ON artist
    BEGIN
        INSERT INTO artist_fts (artist_fts, rowid, name)
        VALUES ('delete', old.id, old.name);
    END;

CREATE TRIGGER artist_au AFTER UPDATE ON artist
    BEGIN
        INSERT INTO artist_fts (artist_fts, rowid, name)
        VALUES ('delete', old.id, old.name);
        INSERT INTO artist_fts (rowid, name)
        VALUES (new.id, new.name);
    END;

CREATE TRIGGER album_ai AFTER INSERT ON album
    BEGIN
        INSERT INTO album_fts (rowid, name)
        VALUES (new.id, new.name);
    END;

CREATE TRIGGER album_ad AFTER DELETE ON album
    BEGIN
        INSERT INTO album_fts (album_fts, rowid, name)
        VALUES ('delete', old.id, old.name);
    END;

CREATE TRIGGER album_au AFTER UPDATE ON album
    BEGIN
        INSERT INTO album_fts (album_fts, rowid, name)
        VALUES ('delete', old.id, old.name);
        INSERT INTO album_fts (rowid, name)
        VALUES (new.id, new.name);
    END;

CREATE TRIGGER song_ai AFTER INSERT ON song
    BEGIN
        INSERT INTO song_fts (rowid, name, lyrics, year)
        VALUES (new.id, new.name, new.lyrics, new.year);
    END;

CREATE TRIGGER song_ad AFTER DELETE ON song
    BEGIN
        INSERT INTO song_fts (song_fts, rowid, name, lyrics, year)
        VALUES ('delete', old.id, old.name, old.lyrics, old.year);
    END;

CREATE TRIGGER song_au AFTER UPDATE ON song
    BEGIN
        INSERT INTO song_fts (song_fts, rowid, name, lyrics, year)
        VALUES ('delete', old.id, old.name, old.lyrics, old.year);
        INSERT INTO song_fts (rowid, name, lyrics, year)
        VALUES (new.id, new.name, new.lyrics, new.year);
    END;

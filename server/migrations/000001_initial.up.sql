CREATE TABLE artist (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR NOT NULL,
    path VARCHAR NOT NULL
);

CREATE TABLE album (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR NOT NULL,
    artist_id INTEGER NOT NULL,
    path VARCHAR NOT NULL,
    album_art VARCHAR,
    FOREIGN KEY(artist_id) REFERENCES artist(id)
);
    
CREATE TABLE song (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    number INTEGER,
    name VARCHAR NOT NULL,
    album_id INTEGER NOT NULL,
    path VARCHAR NOT NULL,
    lyrics VARCHAR,
    file_type VARCHAR NOT NULL,
    year INTEGER,
    FOREIGN KEY(album_id) REFERENCES album(id)
);

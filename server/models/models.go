package models

import "github.com/jmoiron/sqlx"

type Artist struct {
	ID   int64
	Name string
	Path string
}

type Album struct {
	ID       int64
	ArtistID int64 `db:"artist_id"`
	Name     string
	Path     string
	AlbumArt string `db:"album_art"`
}

type Song struct {
	ID      int64
	AlbumID int64 `db:"album_id"`
	Number  int32
	Name    string
	Path    string
}

func SetupTables(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE TABLE artist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		path VARCHAR NOT NULL
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE album (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		artist_id INTEGER NOT NULL,
		path VARCHAR NOT NULL,
		album_art VARCHAR,
		FOREIGN KEY(artist_id) REFERENCES artist(id)
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE song (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		number INTEGER,
		name VARCHAR NOT NULL,
		album_id INTEGER NOT NULL,
		path VARCHAR NOT NULL,
		FOREIGN KEY(album_id) REFERENCES album(id)
	);`)
	if err != nil {
		return err
	}

	return nil
}

package models

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

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
	ID       int64
	AlbumID  int64 `db:"album_id"`
	Number   int32
	Name     string
	Path     string
	Lyrics   string
	FileType string `db:"file_type"`
	Year     int
}

func SetupTables(DBPath string, fs embed.FS) error {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	// driver, err := sqlite.WithInstance(db.DB, &sqlite.Config{})
	// if err != nil {
	// 	return err
	// }
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"iofs:///migrations",
	// 	"sqlite", driver)
	m, err := migrate.NewWithSourceInstance("iofs", d, fmt.Sprintf("sqlite://%s", DBPath))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return err
	}

	return m.Up()
}

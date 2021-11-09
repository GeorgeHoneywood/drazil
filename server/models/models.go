package models

import (
	"embed"
	"errors"
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
	AlbumArt []byte `db:"album_art"`
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

	m, err := migrate.NewWithSourceInstance("iofs", d, fmt.Sprintf("sqlite://%s", DBPath))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return err
	}

	res := m.Up()

	if errors.Is(res, migrate.ErrNoChange) {
		// no change is not an error here

		// TODO: log here
		return nil
	} else if res != nil {
		return res
	} else {
		return res
	}
}

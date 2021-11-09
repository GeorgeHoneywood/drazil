package scanner

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/GeorgeHoneywood/drazil/models"
	"github.com/jmoiron/sqlx"
	"github.com/wader/tag"
	"go.uber.org/zap"
	"golang.org/x/image/draw"
)

func (sc *Scanner) FindArt(db *sqlx.DB) error {
	startTime := time.Now()

	albums := []models.Album{}
	err := db.Select(&albums,
		`SELECT * FROM album
		WHERE album_art IS NULL;`)
	if err != nil {
		sc.Log.Error("could not get albums from db", zap.Error(err))
		return err
	}

	totalBytes := 0

	for _, album := range albums {
		// scan for album artworks here :)
		// try easy ones first, like cover.jpg
		// then we can try using the tags library to retrieve an embedded image

		coverJpg := filepath.Join(sc.MusicPath, album.Path, "cover.jpg")
		if _, err := os.Stat(coverJpg); err == nil {
			fmt.Printf("FOUND\tcoverJpg: %v\n", coverJpg)
			file, _ := os.Open(coverJpg)
			// FIXME: this is a bit redundantW

			orig, _, err := image.Decode(file)
			if err != nil {
				sc.Log.Error("could not decode art file",
					zap.String("path", coverJpg),
				)
				continue
			}

			resized := image.NewRGBA(image.Rect(0, 0, 500, 500))

			draw.NearestNeighbor.Scale(resized, resized.Rect, orig, orig.Bounds(), draw.Over, nil)

			output := bytes.Buffer{}

			jpeg.Encode(&output, resized, nil)

			album.AlbumArt = output.Bytes()

			_, err = db.NamedExec(
				`UPDATE album
				SET album_art = :album_art
				WHERE id = :id;`,
				album)

			if err != nil {
				sc.Log.Error("could not insert art to db",
					zap.String("path", coverJpg),
				)
				continue
			}
			totalBytes += len(album.AlbumArt)
			continue
		}

		song := models.Song{}
		err := db.Get(&song,
			`SELECT * FROM song
			WHERE album_id = $1
			LIMIT 1;`,
			album.ID)

		if err != nil {
			sc.Log.Warn("could not get song from album",
				zap.Error(err),
				zap.String("album_path", album.Path),
			)
			continue
		}

		fmt.Println(song.Path)

		f, err := os.Open(filepath.Join(sc.MusicPath, song.Path))
		if err != nil {
			return err
		}

		meta, err := tag.ReadFrom(f)
		if err != nil {
			sc.Log.Error("could not read tags", zap.String("song_path", song.Path))
			return err
		}

		if meta.Picture() == nil {
			sc.Log.Error("no picture embedded", zap.String("song_path", song.Path))
			continue
		}

		input := bytes.Buffer{}

		input.Write(meta.Picture().Data)

		orig, _, err := image.Decode(&input)
		if err != nil {
			sc.Log.Error("could not decode art file",
				zap.String("path", coverJpg),
			)
			continue
		}

		resized := image.NewRGBA(image.Rect(0, 0, 500, 500))

		draw.NearestNeighbor.Scale(resized, resized.Rect, orig, orig.Bounds(), draw.Over, nil)

		output := bytes.Buffer{}

		jpeg.Encode(&output, resized, nil)

		album.AlbumArt = output.Bytes()

		_, err = db.NamedExec(
			`UPDATE album
			SET album_art = :album_art
			WHERE id = :id;`,
			album)

		if err != nil {
			sc.Log.Error("could not insert art to db",
				zap.String("path", coverJpg),
			)
			continue
		}
		totalBytes += len(album.AlbumArt)

		fmt.Printf("NO\tcoverJpg: %v\n", coverJpg)

	}

	timeTaken := time.Since(startTime)
	sc.Log.Info("finished scanning for art",
		zap.Duration("duration", timeTaken.Round(time.Millisecond)),
		zap.Int("total_bytes_mb", totalBytes/1000/1000),
	)

	return nil
}

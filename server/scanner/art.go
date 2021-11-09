package scanner

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
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
	totalCoverJpg := 0
	totalMetaImage := 0

	for _, album := range albums {
		// scan for album artworks here :)
		// try easy ones first, like cover.jpg
		// then we can try using the tags library to retrieve an embedded image

		coverJpg := filepath.Join(sc.MusicPath, album.Path, "cover.jpg")
		if file, err := os.Open(coverJpg); err == nil {
			output, err := sc.scaleImage(file)
			if err != nil {
				sc.Log.Error("could not scale image",
					zap.String("image_path", coverJpg),
					zap.Error(err),
				)
				continue
			}
			album.AlbumArt = output.Bytes()
			totalCoverJpg += 1
		} else {
			song := models.Song{}
			err := db.Get(&song,
				`SELECT * FROM song
				WHERE album_id = $1
				LIMIT 1;`,
				album.ID)

			if err != nil {
				// sc.Log.Warn("could not get song from album",
				// 	zap.Error(err),
				// 	zap.String("album_path", album.Path),
				// )
				continue
			}

			f, err := os.Open(filepath.Join(sc.MusicPath, song.Path))
			if err != nil {
				sc.Log.Error("could not open song", zap.String("song_path", song.Path))
				continue
			}

			meta, err := tag.ReadFrom(f)
			if err != nil {
				sc.Log.Error("could not read tags", zap.String("song_path", song.Path))
				continue
			}

			if meta.Picture() == nil {
				// sc.Log.Error("no picture embedded", zap.String("song_path", song.Path))
				continue
			}

			output, err := sc.scaleImage(bytes.NewReader(meta.Picture().Data))
			if err != nil {
				sc.Log.Error("could not scale image",
					zap.String("song", song.Path),
					zap.Error(err),
				)
				continue
			}

			album.AlbumArt = output.Bytes()
			totalMetaImage += 1
		}

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
	}

	timeTaken := time.Since(startTime)
	sc.Log.Info("finished scanning for art",
		zap.Duration("duration", timeTaken.Round(time.Millisecond)),
		zap.Int("total_bytes_mb", totalBytes/1000/1000),
		zap.Int("total_cover_jpg", totalCoverJpg),
		zap.Int("total_meta_image", totalMetaImage),
	)

	return nil
}

func (sc *Scanner) scaleImage(src io.Reader) (*bytes.Buffer, error) {
	orig, _, err := image.Decode(src)
	if err != nil {
		return nil, fmt.Errorf("could not decode art file: %s", err)
	}

	resized := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.NearestNeighbor.Scale(resized, resized.Rect, orig, orig.Bounds(), draw.Over, nil)

	output := bytes.Buffer{}
	jpeg.Encode(&output, resized, nil)

	return &output, nil
}

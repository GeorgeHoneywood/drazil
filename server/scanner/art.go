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
	"sync"
	"time"

	"github.com/GeorgeHoneywood/drazil/models"
	"github.com/jmoiron/sqlx"
	"github.com/wader/tag"
	"go.uber.org/zap"
	"golang.org/x/image/draw"
)

type ArtStats struct {
	totalBytes     int
	totalCoverJpg  int
	totalMetaImage int
}

func (sc *Scanner) FindArt(db *sqlx.DB) error {
	parallelCount := 4

	startTime := time.Now()
	s := &ArtStats{}

	albums := []models.Album{}
	err := db.Select(&albums,
		`SELECT * FROM album
		WHERE album_art IS NULL;`)
	if err != nil {
		sc.Log.Error("could not get albums from db", zap.Error(err))
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	c := make(chan models.Album)

	var wg sync.WaitGroup
	wg.Add(parallelCount)
	for ii := 0; ii < parallelCount; ii++ {

		go func(c chan models.Album) {
			for {
				album, more := <-c
				if !more {
					wg.Done()
					return
				}

				sc.processImage(album, s, tx)
			}
		}(c)
	}

	for _, album := range albums {
		c <- album
	}
	close(c)
	wg.Wait()

	err = tx.Commit()
	if err != nil {
		return err
	}

	timeTaken := time.Since(startTime)
	sc.Log.Info("finished scanning for art",
		zap.Duration("duration", timeTaken.Round(time.Millisecond)),
		zap.Int("total_bytes_mb", s.totalBytes/1000/1000),
		zap.Int("total_cover_jpg", s.totalCoverJpg),
		zap.Int("total_meta_image", s.totalMetaImage),
	)

	return nil
}

func (sc *Scanner) processImage(album models.Album, s *ArtStats, tx *sqlx.Tx) {
	// scan for album artworks here :)
	// try easy ones first, like cover.jpg
	// then we can try using the tags library to retrieve an embedded image

	coverJpg := filepath.Join(sc.MusicPath, album.Path, "cover.jpg")
	if file, err := os.Open(coverJpg); err == nil {
		output, err := sc.scaleImage(file)
		if err != nil {
			// sc.Log.Warn("could not scale image",
			// 	zap.String("image_path", coverJpg),
			// 	zap.Error(err),
			// )
			return
		}
		album.AlbumArt = output.Bytes()
		s.totalCoverJpg++
	} else {
		song := models.Song{}

		err := tx.Get(&song,
			`SELECT * FROM song
			WHERE album_id = $1
			LIMIT 1;`,
			album.ID)
		if err != nil {
			sc.Log.Warn("could not get song from album",
				zap.Error(err),
				zap.String("album_path", album.Path),
			)
			return
		}

		f, err := os.Open(filepath.Join(sc.MusicPath, song.Path))
		if err != nil {
			sc.Log.Error("could not open song", zap.String("song_path", song.Path))
			return
		}

		meta, err := tag.ReadFrom(f)
		if err != nil {
			sc.Log.Error("could not read tags", zap.String("song_path", song.Path))
			return
		}

		if meta.Picture() == nil {
			// sc.Log.Error("no picture embedded", zap.String("song_path", song.Path))
			return
		}

		output, err := sc.scaleImage(bytes.NewReader(meta.Picture().Data))
		if err != nil {
			// sc.Log.Warn("could not scale image",
			// 	zap.String("song", song.Path),
			// 	zap.Error(err),
			// )
			return
		}

		album.AlbumArt = output.Bytes()
		s.totalMetaImage++
	}

	_, err := tx.NamedExec(
		`UPDATE album
		SET album_art = :album_art
		WHERE id = :id;`,
		album)

	if err != nil {
		sc.Log.Error("could not insert art to db",
			zap.String("path", coverJpg),
			zap.Error(err),
		)
		return
	}

	s.totalBytes += len(album.AlbumArt)
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

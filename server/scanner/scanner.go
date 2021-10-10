package scanner

import (
	"database/sql"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/JoeRourke123/Monkey/models"
	"github.com/jmoiron/sqlx"
	"github.com/wader/tag"
	"go.uber.org/zap"
)

type Stats struct {
	NewArtists      int
	ExistingArtists int
	NewAlbums       int
	ExistingAlbums  int
	NewSongs        int
	ExistingSongs   int
}

type Scanner struct {
	MusicPath string
	Log       *zap.Logger
	s         Stats
}

func (sc *Scanner) FindArtists(db *sqlx.DB) error {
	startTime := time.Now()

	artist_names, err := os.ReadDir(sc.MusicPath)
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, artist_name := range artist_names {
		if !artist_name.IsDir() ||
			strings.HasPrefix(artist_name.Name(), ".") {
			continue
		}

		artist := &models.Artist{
			Name: artist_name.Name(),
			Path: artist_name.Name(),
		}

		err := tx.Get(artist, `
		SELECT * 
		FROM artist
        WHERE path=$1`, artist.Path)

		if err == sql.ErrNoRows {
			rows, err := tx.NamedQuery(`
			INSERT OR REPLACE INTO artist (name, path)
			VALUES (:name, :path)
			RETURNING id`, artist)
			if err != nil {
				return err
			}

			if rows.Next() {
				rows.Scan(&artist.ID)
			}

			sc.s.NewArtists++
		} else if err != nil {
			return err
		} else {
			sc.s.ExistingArtists++
		}

		err = sc.findAlbums(tx, artist)
		if err != nil {
			return err
		}

	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	timeTaken := time.Since(startTime)
	sc.Log.Info("finished scanning for music",
		zap.Duration("duration", timeTaken.Round(time.Millisecond)),
		zap.Int("new_artists", sc.s.NewArtists),
		zap.Int("existing_artists", sc.s.ExistingArtists),
		zap.Int("new_albums", sc.s.NewAlbums),
		zap.Int("existing_albums", sc.s.ExistingAlbums),
		zap.Int("new_songs", sc.s.NewSongs),
		zap.Int("existing_songs", sc.s.ExistingSongs),
	)

	return nil
}

func (sc *Scanner) findAlbums(tx *sqlx.Tx, artist *models.Artist) error {
	album_names, err := os.ReadDir(sc.MusicPath + artist.Path)
	if err != nil {
		return err
	}

	for _, album_name := range album_names {
		if !album_name.IsDir() ||
			strings.HasPrefix(album_name.Name(), ".") {
			continue
		}

		path := strings.Builder{}
		path.WriteString(artist.Path)
		path.WriteString("/")
		path.WriteString(album_name.Name())

		album := &models.Album{
			ArtistID: artist.ID,
			Path:     path.String(),
		}

		err := tx.Get(album, `
		SELECT * 
		FROM album
        WHERE path=$1 AND artist_id=$2`, album.Path, artist.ID)

		if errors.Is(err, sql.ErrNoRows) {
			album.Name = album_name.Name()
			rm := regexp.MustCompile(`\(.*\)`)
			album.Name = rm.ReplaceAllString(album.Name, "")
			album.Name = strings.Trim(album.Name, "- ")
			if strings.HasPrefix(album.Name, artist.Name) &&
				!strings.HasSuffix(album.Name, artist.Name) {
				album.Name = strings.Replace(album.Name, artist.Name, "", -1)
			}
			album.Name = strings.Trim(album.Name, "- ")

			path.WriteString("/cover.jpg")

			if _, err := os.Stat(sc.MusicPath + path.String()); err == nil {
				album.AlbumArt = path.String()
			}

			rows, err := tx.NamedQuery(`
			INSERT INTO album (artist_id, name, path, album_art)
			VALUES (:artist_id, :name, :path, :album_art)
			RETURNING id;`, album)
			if err != nil {
				return err
			}
			if rows.Next() {
				rows.Scan(&album.ID)
			}

			sc.s.NewAlbums++
		} else if err != nil {
			return err
		} else {
			sc.s.ExistingAlbums++
		}

		err = sc.findSongs(tx, album, artist)
		if err != nil {
			return err
		}

	}

	return nil
}

func (sc *Scanner) findSongs(tx *sqlx.Tx, album *models.Album, artist *models.Artist) error {
	song_names, err := os.ReadDir(sc.MusicPath + album.Path)
	if err != nil {
		return err
	}

	for i, song_name := range song_names {
		if song_name.IsDir() {
			// this must be a Disc 01/CD 01 folder or so
			// TODO: remember which disc each song is on
			a := *album
			builder := strings.Builder{}
			builder.WriteString(album.Path)
			builder.WriteString("/")
			builder.WriteString(song_name.Name())
			a.Path = builder.String()

			err := sc.findSongs(tx, &a, artist)
			if err != nil {
				return err
			}
			continue
		}
		if strings.HasPrefix(song_name.Name(), ".") {
			continue
		}

		if !(strings.HasSuffix(song_name.Name(), ".mp3") ||
			strings.HasSuffix(song_name.Name(), ".ogg") ||
			strings.HasSuffix(song_name.Name(), ".opus") ||
			strings.HasSuffix(song_name.Name(), ".flac")) {
			continue
		}

		path := strings.Builder{}
		path.WriteString(album.Path)
		path.WriteString("/")
		path.WriteString(song_name.Name())

		song := &models.Song{
			AlbumID: album.ID,
			Path:    path.String(),
		}

		err := tx.Get(song, `
		SELECT * 
		FROM song
        WHERE path=$1 AND album_id=$2`, song.Path, album.ID)

		if errors.Is(err, sql.ErrNoRows) {
			rt := regexp.MustCompile(`^\d+`)
			trackNumberString := rt.FindString(song_name.Name())

			trackName := song_name.Name()
			if trackNumberString != "" {
				split := strings.Split(trackName, ".")
				trackName = strings.Join(split[:len(split)-1], ".")

				rm := regexp.MustCompile(`\[.*\]`)
				trackName = rm.ReplaceAllString(trackName, "")
				trackName = strings.Trim(trackName, "-. ")

				// this handles cases where the track name is the same as the album/artist name
				// i.e. if trackName = `AC/DC - Highway to Hell`, we want to make sure we remove
				// the AC/DC, but not `Highway to Hell`, even though that is the album name
				if !strings.HasSuffix(trackName, trackNumberString) {
					trackName = strings.Replace(trackName, trackNumberString, "", -1)
				}
				if !strings.HasSuffix(trackName, album.Name) {
					trackName = strings.Replace(trackName, album.Name, "", -1)
				}
				if !strings.HasSuffix(trackName, artist.Name) {
					trackName = strings.Replace(trackName, artist.Name, "", -1)
				}

				trackName = strings.Trim(trackName, "-. ")
			}
			song.Name = trackName

			// n.b. this variable exists because there is no ParseInt32 function
			// don't ask me why...
			songNumber, err := strconv.Atoi(trackNumberString)
			if err != nil {
				// if we can't parse a track number from the name we can just make a guess
				// TODO: should attempt to use the song tags first
				songNumber = i + 1
			}

			song.Number = int32(songNumber)

			f, err := os.Open(sc.MusicPath + path.String())
			if err != nil {
				return err
			}

			meta, err := tag.ReadFrom(f)
			if err != nil {
				sc.Log.Error("could not read tags", zap.String("song_path", song.Path))
				return err
			}

			if meta != nil {
				song.Lyrics = meta.Lyrics()
				song.FileType = string(meta.FileType())
				song.Year = meta.Year()
			}

			_, err = tx.NamedExec(`
			INSERT INTO song (album_id, number, name, path, lyrics, file_type, year)
			VALUES (:album_id, :number, :name, :path, :lyrics, :file_type, :year)
			RETURNING id;`, song)
			if err != nil {
				return err
			}

			sc.s.NewSongs++
		} else if err != nil {
			return err
		} else {
			sc.s.ExistingSongs++
		}
	}

	return nil
}

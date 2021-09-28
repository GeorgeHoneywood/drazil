package scanner

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/JoeRourke123/Monkey/models"
	"github.com/jmoiron/sqlx"
	"github.com/wader/tag"
)

func FindArtists(root string, db *sqlx.DB) error {
	startTime := time.Now()

	artist_names, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, artist_name := range artist_names {
		if !artist_name.IsDir() || strings.HasPrefix(artist_name.Name(), ".") {
			continue
		}

		artist := &models.Artist{
			Name: artist_name.Name(),
			Path: artist_name.Name(),
		}

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

		err = findAlbums(tx, root, artist)
		if err != nil {
			return err
		}

	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	timeTaken := time.Since(startTime)
	fmt.Printf("endTime: %s\n", timeTaken.Round(time.Millisecond).String())

	return nil
}

func findAlbums(tx *sqlx.Tx, root string, artist *models.Artist) error {
	album_names, err := os.ReadDir(root + artist.Path)
	if err != nil {
		return err
	}

	for _, album_name := range album_names {
		if !album_name.IsDir() || strings.HasPrefix(album_name.Name(), ".") {
			continue
		}

		albumName := album_name.Name()

		rm := regexp.MustCompile(`\(.*\)`)
		albumName = rm.ReplaceAllString(albumName, "")

		albumName = strings.Trim(albumName, "- ")

		if strings.HasPrefix(albumName, artist.Name) && !strings.HasSuffix(albumName, artist.Name) {
			albumName = strings.Replace(albumName, artist.Name, "", -1)
		}

		albumName = strings.Trim(albumName, "- ")

		path := strings.Builder{}
		path.WriteString(artist.Path)
		path.WriteString("/")
		path.WriteString(album_name.Name())

		album := &models.Album{
			Name:     albumName,
			ArtistID: artist.ID,
			Path:     path.String(),
		}

		path.WriteString("/cover.jpg")

		if _, err := os.Stat(root + path.String()); err == nil {
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

		err = findSongs(tx, root, album, artist)
		if err != nil {
			return err
		}

	}

	return nil
}

func findSongs(tx *sqlx.Tx, root string, album *models.Album, artist *models.Artist) error {
	song_names, err := os.ReadDir(root + album.Path)
	if err != nil {
		return err
	}

	for i, song_name := range song_names {
		if song_name.IsDir() {
			// // this must be a Disc 01/CD 01 folder or so
			// builder := strings.Builder{}
			// builder.WriteString(album.Path)
			// builder.WriteString("/")
			// builder.WriteString(song_name.Name())
			// album.Path = builder.String()

			// err := findSongs(tx, album, artist)
			// if err != nil {
			// 	return err
			// }
			continue
		}
		if strings.HasPrefix(song_name.Name(), ".") {
			continue
		}

		if !(strings.HasSuffix(song_name.Name(), ".mp3") || strings.HasSuffix(song_name.Name(), ".ogg")) {
			continue
		}

		rt := regexp.MustCompile(`^\d+`) // \.?
		trackNumber := rt.FindString(song_name.Name())

		trackName := song_name.Name()
		if trackNumber != "" {
			split := strings.Split(trackName, ".")
			trackName = strings.Join(split[:len(split)-1], ".")

			rm := regexp.MustCompile(`\[.*\]`)
			trackName = rm.ReplaceAllString(trackName, "")
			trackName = strings.Trim(trackName, "-. ")

			// this handles cases where the track name is the same as the album/artist name
			if !strings.HasSuffix(trackName, trackNumber) {
				trackName = strings.Replace(trackName, trackNumber, "", -1)
			}
			if !strings.HasSuffix(trackName, album.Name) {
				trackName = strings.Replace(trackName, album.Name, "", -1)
			}
			if !strings.HasSuffix(trackName, artist.Name) {
				trackName = strings.Replace(trackName, album.Name, "", -1)
			}

			trackName = strings.Trim(trackName, "-. ")
		}

		path := strings.Builder{}
		path.WriteString(album.Path)
		path.WriteString("/")
		path.WriteString(song_name.Name())

		trackNumberInt, err := strconv.Atoi(trackNumber)
		if err != nil {
			trackNumberInt = i + 1
		}

		f, err := os.Open(root + path.String())

		if err != nil {
			panic(err)
		}

		meta, err := tag.ReadFrom(f)

		if err != nil {
			fmt.Println(err)
		}

		song := &models.Song{
			AlbumID: album.ID,
			Number:  int32(trackNumberInt),
			Name:    trackName,
			Path:    path.String(),
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

	}

	return nil
}

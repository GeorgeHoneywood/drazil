package scanner

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func (sc *Scanner) SetTotals(db *sqlx.DB) error {
	_, err := db.Exec(`UPDATE album
	SET song_count = sum.count
	FROM (
		SELECT album_id, COUNT(*) AS count
		FROM song
		GROUP BY album_id
		) AS sum
	WHERE sum.album_id = album.id;`)
	if err != nil {
		sc.Log.Error("could not set album song totals", zap.Error(err))
	}

	_, err = db.Exec(`UPDATE artist
	SET album_count = sum.count
	FROM (
		SELECT artist_id, COUNT(*) AS count
		FROM album
		GROUP BY artist_id
		) AS sum
	WHERE sum.artist_id = artist.id;`)
	if err != nil {
		sc.Log.Error("could not set artist album totals", zap.Error(err))
	}

	_, err = db.Exec(`UPDATE artist
	SET song_count = sum.count
	FROM (
		SELECT artist_id, SUM(song_count) AS count
		FROM album
		GROUP BY artist_id
		) AS sum
	WHERE sum.artist_id = artist.id;`)
	if err != nil {
		sc.Log.Error("could not set artist song totals", zap.Error(err))
	}
	return nil
}

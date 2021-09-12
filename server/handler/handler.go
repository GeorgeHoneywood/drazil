package handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/JoeRourke123/Monkey/models"
	"github.com/JoeRourke123/Monkey/spec"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	spec.UnimplementedMonkeyServer
	DB        *sqlx.DB
	HTTPRoot  string
	MusicPath string
}

func (h *Handler) ListArtists(ctx context.Context, in *spec.ArtistsRequest) (*spec.ArtistsReply, error) {
	artists := []models.Artist{}
	err := h.DB.Select(&artists,
		`SELECT * FROM artist;`)
	if err != nil {
		log.Print(err)
	}

	var artistCount int32
	err = h.DB.Get(&artistCount,
		`SELECT COUNT(*) FROM artist;`)
	if err != nil {
		log.Print(err)
	}

	out := make([]*spec.Artist, len(artists))
	for i, artist := range artists {
		out[i] = &spec.Artist{
			Name:     artist.Name,
			ImageUrl: "yzx",
			Id:       artist.ID,
		}
	}

	return &spec.ArtistsReply{
		Artists:     out,
		ArtistCount: artistCount,
	}, nil
}

func (h *Handler) ListAlbums(ctx context.Context, in *spec.AlbumsRequest) (*spec.AlbumsReply, error) {

	albums := []models.Album{}
	err := h.DB.Select(&albums,
		`SELECT * FROM album 
		WHERE artist_id=$1;`,
		in.ArtistId)
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if len(albums) < 1 {
		return nil, status.Error(codes.NotFound, "No albums found for this artist")
	}

	artistName := ""
	err = h.DB.Get(&artistName,
		`SELECT name FROM artist 
		WHERE id=$1;`,
		in.ArtistId)
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := make([]*spec.Album, len(albums))
	for i, album := range albums {
		out[i] = &spec.Album{
			Id:   album.ID,
			Name: album.Name,
		}

		if album.AlbumArt != "" {
			out[i].AlbumArt = h.HTTPRoot + "/media/" + album.AlbumArt
		}
	}

	return &spec.AlbumsReply{
		Albums:     out,
		ArtistName: artistName,
	}, nil
}

func (h *Handler) ListSongs(ctx context.Context, in *spec.SongsRequest) (*spec.SongsReply, error) {
	album := &models.Album{}
	err := h.DB.Get(album,
		`SELECT * FROM album 
		WHERE id=$1 AND artist_id=$2`,
		in.AlbumId,
		in.ArtistId)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("artist '%d' has no album '%d'", in.ArtistId, in.AlbumId))
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	songs := []models.Song{}
	err = h.DB.Select(&songs,
		`SELECT * FROM song
		WHERE album_id=$1
		ORDER BY number;`,
		in.AlbumId)
	if err != nil {
		log.Print(err)
	}

	artistName := ""
	err = h.DB.Get(&artistName,
		`SELECT name FROM artist 
		WHERE id=$1;`,
		in.ArtistId)
	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err != nil {
		log.Print(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := make([]*spec.Song, len(songs))
	for i, song := range songs {
		out[i] = &spec.Song{
			Name:   song.Name,
			Number: song.Number,
			Path:   h.HTTPRoot + "/media/" + song.Path,
		}
	}

	return &spec.SongsReply{
		Songs:      out,
		ArtistName: artistName,
		AlbumName:  album.Name,
		AlbumArt:   h.HTTPRoot + "/media/" + album.AlbumArt,
	}, nil
}

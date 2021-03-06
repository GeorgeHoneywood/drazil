package handler

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/GeorgeHoneywood/drazil/models"
	"github.com/GeorgeHoneywood/drazil/spec"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	spec.UnimplementedDrazilServer
	DB        *sqlx.DB
	MusicPath string
	Log       *zap.Logger
}

func (h *Handler) ListArtists(ctx context.Context, in *spec.ArtistsRequest) (*spec.ArtistsReply, error) {
	artists := []models.Artist{}
	err := h.DB.Select(&artists,
		`SELECT * FROM artist ORDER BY name;`)
	if err != nil {
		h.Log.Error("could not get artists from db", zap.Error(err))
	}

	var artistCount int32
	err = h.DB.Get(&artistCount,
		`SELECT COUNT(*) FROM artist;`)
	if err != nil {
		h.Log.Error("could not get artist count from db", zap.Error(err))
	}

	out := make([]*spec.Artist, len(artists))
	for i, artist := range artists {
		out[i] = &spec.Artist{
			Name:       artist.Name,
			ImageUrl:   "yzx",
			Id:         artist.ID,
			AlbumCount: artist.AlbumCount,
			SongCount:  artist.SongCount,
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
		WHERE artist_id=$1
		ORDER BY name;`,
		in.ArtistId)
	if err != nil {
		h.Log.Error("could not get albums from db", zap.Error(err))
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
		h.Log.Error("could not get artist from db", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := make([]*spec.Album, len(albums))
	for i, album := range albums {
		out[i] = &spec.Album{
			Id:        album.ID,
			Name:      album.Name,
			SongCount: album.SongCount,
			Year:      int64(album.Year),
		}
	}

	return &spec.AlbumsReply{
		Albums:     out,
		ArtistName: artistName,
		ArtistId:   in.ArtistId,
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
		h.Log.Error("could not get album from db", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	songs := []models.Song{}
	err = h.DB.Select(&songs,
		`SELECT * FROM song
		WHERE album_id=$1
		ORDER BY number;`,
		in.AlbumId)
	if err != nil {
		h.Log.Error("could not get songs from db", zap.Error(err), zap.Int64("album_id", in.AlbumId))
		return nil, status.Error(codes.Internal, err.Error())
	}

	artistName := ""
	err = h.DB.Get(&artistName,
		`SELECT name FROM artist 
		WHERE id=$1;`,
		in.ArtistId)
	if err != nil {
		h.Log.Error("could not get artist from db", zap.Error(err), zap.Int64("artist_id", in.ArtistId))
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := make([]*spec.Song, len(songs))
	for i, song := range songs {
		out[i] = &spec.Song{
			Id:       song.ID,
			Name:     song.Name,
			Number:   song.Number,
			Lyrics:   song.Lyrics,
			FileType: song.FileType,
		}
	}

	return &spec.SongsReply{
		Songs:      out,
		ArtistName: artistName,
		AlbumName:  album.Name,
	}, nil
}

func (h *Handler) ListAllAlbums(ctx context.Context, in *spec.AllAlbumsRequest) (*spec.AllAlbumsReply, error) {
	albums := []models.Album{}
	err := h.DB.Select(&albums,
		`SELECT * FROM album ORDER BY name;`)
	if err != nil {
		h.Log.Error("could not get albums from db", zap.Error(err))
	}

	out := make([]*spec.AllAlbum, len(albums))
	for i, album := range albums {
		out[i] = &spec.AllAlbum{
			Id:         album.ID,
			Name:       album.Name,
			ArtistName: "Lorem Ipsum",
			ArtistId:   album.ArtistID,
		}
	}

	return &spec.AllAlbumsReply{
		Albums: out,
	}, nil
}

func (h *Handler) Search(ctx context.Context, in *spec.SearchRequest) (*spec.SearchReply, error) {
	artists := []models.Artist{}
	err := h.DB.Select(&artists,
		`SELECT * FROM artist_fts($1)
		JOIN artist ON artist.id = artist_fts.rowid
		ORDER BY rank;`,
		in.Token)
	if err != nil {
		h.Log.Error("could not get artists from db", zap.Error(err))
	}

	outArtists := make([]*spec.Artist, len(artists))
	for i, artist := range artists {
		outArtists[i] = &spec.Artist{
			Name: artist.Name,
			Id:   artist.ID,
		}
	}

	albums := []models.Album{}
	err = h.DB.Select(&albums,
		`SELECT * FROM album_fts($1)
		JOIN album ON album.id = album_fts.rowid
		ORDER BY rank;`,
		in.Token)
	if err != nil {
		h.Log.Error("could not get albums from db", zap.Error(err))
	}

	outAlbums := make([]*spec.AllAlbum, len(albums))
	for i, album := range albums {
		outAlbums[i] = &spec.AllAlbum{
			Id:         album.ID,
			Name:       album.Name,
			ArtistName: "Lorem Ipsum",
			ArtistId:   album.ArtistID,
		}
	}

	songs := []models.Song{}
	err = h.DB.Select(&songs,
		`SELECT * FROM song_fts($1)
		JOIN song ON song.id = song_fts.rowid
		ORDER BY rank;`,
		in.Token)
	if err != nil {
		h.Log.Error("could not get songs from db", zap.Error(err))
	}

	outSongs := make([]*spec.AllSong, len(songs))
	for i, song := range songs {
		outSongs[i] = &spec.AllSong{
			AlbumId:    song.AlbumID,
			AlbumName:  "",
			ArtistName: "",
			ArtistId:   0,
			Id:         song.ID,
			Name:       song.Name,
			Number:     song.Number,
			Lyrics:     song.Lyrics,
			FileType:   song.FileType,
		}
	}

	return &spec.SearchReply{
		Artists: outArtists,
		Albums:  outAlbums,
		Songs:   outSongs,
	}, nil
}

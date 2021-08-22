/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"context"
	"flag"
	"net/http"

	pb "github.com/JoeRourke123/Monkey/spec"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// gw "github.com/yourorg/yourrepo/proto/gen/go/your/service/v1/your_service"
)

// Update
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedMonkeyServer
	//db *sql.DB
}

const (
	port         = "localhost:9090"
	musicPath    = "/home/honeyfox/Music/"
	databasePath = "./monkey.db"
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
}
type Song struct {
	ID      int64
	AlbumID int64 `db:"album_id"`
	Number  int32
	Name    string
	Path    string
}

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
	db                 *sqlx.DB
)

func (s *server) ListArtists(ctx context.Context, in *pb.ArtistsRequest) (*pb.ArtistsReply, error) {
	artists := []Artist{}
	err := db.Select(&artists,
		`SELECT * FROM artist;`)
	if err != nil {
		log.Print(err)
	}

	out := make([]*pb.Artist, len(artists))
	for i, artist := range artists {
		out[i] = &pb.Artist{
			Name:     artist.Name,
			ImageUrl: "yzx",
			Id:       artist.ID,
		}
	}

	return &pb.ArtistsReply{
		Artists: out,
	}, nil
}

func (s *server) ListAlbums(ctx context.Context, in *pb.AlbumsRequest) (*pb.AlbumsReply, error) {

	albums := []Album{}
	err := db.Select(&albums,
		`SELECT * FROM album 
		WHERE artist_id=$1;`,
		in.ArtistId)
	if err != nil {
		log.Print(err)
	}

	out := make([]*pb.Album, len(albums))
	for i, album := range albums {
		// 	songs := []Song{}
		// 	err := db.Select(&songs,
		// 		`SELECT * FROM song
		// 		WHERE album_id=$1;`,
		// 		album.ID)
		// 	if err != nil {
		// 		log.Print(err)
		// 	}

		// 	songsOut := make([]*pb.Song, len(songs))
		// 	for j, song := range songs {
		// 		songsOut[j] = &pb.Song{
		// 			Name:   song.Name,
		// 			Number: song.Number,
		// 		}
		// 	}

		out[i] = &pb.Album{
			Id:       album.ID,
			ImageUrl: "xyz",
			Name:     album.Name,
		}
	}

	return &pb.AlbumsReply{
		Albums: out,
	}, nil
}

func (s *server) ListSongs(ctx context.Context, in *pb.SongsRequest) (*pb.SongsReply, error) {
	album := &Album{}
	err := db.Get(album,
		`SELECT * FROM album 
		WHERE id=$1 AND artist_id=$2`,
		in.AlbumId,
		in.ArtistId)

	if err == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("artist '%d' has no album '%d'", in.ArtistId, in.AlbumId))
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	songs := []Song{}
	err = db.Select(&songs,
		`SELECT * FROM song
		WHERE album_id=$1
		ORDER BY number;`,
		in.AlbumId)
	if err != nil {
		log.Print(err)
	}

	out := make([]*pb.Song, len(songs))
	for i, song := range songs {
		out[i] = &pb.Song{
			Name:   song.Name,
			Number: song.Number,
		}
	}

	return &pb.SongsReply{
		Songs: out,
	}, nil
}

func findArtists(dir string) error {
	startTime := time.Now()

	artist_names, err := os.ReadDir(dir)
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

		path := strings.Builder{}
		path.WriteString(dir)
		path.WriteString(artist_name.Name())

		artist := &Artist{
			Name: artist_name.Name(),
			Path: path.String(),
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

		err = findAlbums(tx, path.String(), artist_name.Name(), artist)
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

func findAlbums(tx *sqlx.Tx, dir string, artist_name string, artist *Artist) error {
	album_names, err := os.ReadDir(dir)
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

		if strings.HasPrefix(albumName, artist_name) && !strings.HasSuffix(albumName, artist_name) {
			albumName = strings.Replace(albumName, artist_name, "", -1)
		}

		albumName = strings.Trim(albumName, "- ")

		path := strings.Builder{}
		path.WriteString(dir)
		path.WriteString("/")
		path.WriteString(album_name.Name())

		album := &Album{
			Name:     albumName,
			ArtistID: artist.ID,
			Path:     path.String(),
		}

		rows, err := tx.NamedQuery(`
		INSERT INTO album (artist_id, name, path)
        VALUES (:artist_id, :name, :path)
		RETURNING id;`, album)
		if err != nil {
			return err
		}
		if rows.Next() {
			rows.Scan(&album.ID)
		}

		err = findSongs(tx, path.String(), artist_name, albumName, album)
		if err != nil {
			return err
		}

	}

	return nil
}

func findSongs(tx *sqlx.Tx, dir string, artist_name string, album_name string, album *Album) error {
	song_names, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, song_name := range song_names {
		if song_name.IsDir() {
			// this must be a Disc 01/CD 01 folder or so
			builder := strings.Builder{}
			builder.WriteString(dir)
			builder.WriteString("/")
			builder.WriteString(song_name.Name())

			err := findSongs(tx, builder.String(), artist_name, album_name, album)
			if err != nil {
				return err
			}
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
			if !strings.HasSuffix(trackName, artist_name) {
				trackName = strings.Replace(trackName, artist_name, "", -1)
			}

			trackName = strings.Trim(trackName, "-. ")
		}

		path := strings.Builder{}
		path.WriteString(dir)
		path.WriteString("/")
		path.WriteString(song_name.Name())

		trackNumberInt, err := strconv.Atoi(trackNumber)
		if err != nil {
			trackNumberInt = -1
		}

		song := &Song{
			Number:  int32(trackNumberInt),
			Name:    trackName,
			AlbumID: album.ID,
			Path:    path.String(),
		}

		_, err = tx.NamedExec(`
		INSERT INTO song (album_id, number, name, path)
        VALUES (:album_id, :number, :name, :path)
		RETURNING id;`, song)
		if err != nil {
			return err
		}

	}

	return nil
}

func run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterMonkeyServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = pb.RegisterMonkeyHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	os.Remove(databasePath)

	_, err := os.Stat(databasePath)
	exist := !errors.Is(err, os.ErrNotExist)

	db, err = sqlx.Connect("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if !exist {
		log.Print("setting up tables")

		err = setupTables(db)
		if err != nil {
			log.Fatalf("error setting up db: %s", err.Error())
		}
	}

	err = findArtists(musicPath)
	if err != nil {
		log.Printf("error walking directories: %s", err.Error())
	}

	// artists := []Artist{}
	// err = db.Select(&artists, `SELECT * FROM artist;`)
	// if err != nil {
	// 	log.Print(err)
	// }

	// for i, a := range artists {
	// 	fmt.Printf("i: %v, a: %v\n", i, a)
	// }

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func setupTables(db *sqlx.DB) error {
	_, err := db.Exec(`CREATE TABLE artist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR,
		path VARCHAR
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE album (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR,
		artist_id INTEGER,
		path VARCHAR,
		FOREIGN KEY(artist_id) REFERENCES artist(id)
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE song (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		number INTEGER,
		name VARCHAR,
		album_id INTEGER,
		path VARCHAR,
		FOREIGN KEY(album_id) REFERENCES album(id)
	);`)
	if err != nil {
		return err
	}

	return nil
}

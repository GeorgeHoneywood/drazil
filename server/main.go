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
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"context"
	"flag"
	"net/http"

	pb "github.com/JoeRourke123/Monkey/spec"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	// gw "github.com/yourorg/yourrepo/proto/gen/go/your/service/v1/your_service"
)

// Update
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedMonkeyServer
}

const (
	port      = "localhost:9090"
	musicPath = "/home/honeyfox/Music/"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

// SayHello implements helloworld.GreeterServer
func (s *server) GetAlbums(ctx context.Context, in *pb.AlbumRequest) (*pb.AlbumReply, error) {
	return &pb.AlbumReply{
		Albums: []*pb.Album{{
			ImageUrl:   "eggs",
			AlbumTitle: "ABBA",
			Songs: []*pb.Song{{
				SongTitle:   in.GetName(),
				TrackNumber: "1",
				SongUrl:     "2",
			}},
		}},
	}, nil
}

func findArtists(dir string) error {
	artists, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, artist := range artists {
		if !artist.IsDir() || strings.HasPrefix(artist.Name(), ".") {
			continue
		}
		fmt.Printf("artist.Name(): %v\n", artist.Name())

		builder := strings.Builder{}
		builder.WriteString(dir)
		builder.WriteString(artist.Name())

		err := findAlbums(builder.String(), artist.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func findAlbums(dir string, artist string) error {
	albums, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, album := range albums {
		if !album.IsDir() || strings.HasPrefix(album.Name(), ".") {
			continue
		}

		albumName := album.Name()
		if strings.HasPrefix(album.Name(), artist) {
			albumName = strings.Replace(albumName, artist, "", -1)
		}

		albumName = strings.Trim(albumName, "- ")
		fmt.Printf("\talbumName: %v\n", albumName)

		builder := strings.Builder{}
		builder.WriteString(dir)
		builder.WriteString("/")
		builder.WriteString(album.Name())

		err := findSongs(builder.String(), artist, albumName)
		if err != nil {
			return err
		}
	}

	return nil
}

func findSongs(dir string, artist string, album string) error {
	songs, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, song := range songs {
		if song.IsDir() {
			// this must be a Disc 01/CD 01 folder or so
			builder := strings.Builder{}
			builder.WriteString(dir)
			builder.WriteString("/")
			builder.WriteString(song.Name())

			err := findSongs(builder.String(), artist, album)
			if err != nil {
				return err
			}
			continue
		}
		if strings.HasPrefix(song.Name(), ".") {
			continue
		}

		r := regexp.MustCompile(`^\d+`) // \.?
		trackNumber := r.FindString(song.Name())

		songName := song.Name()
		if trackNumber != "" {
			songName = strings.Replace(songName, trackNumber, "", -1)
			songName = strings.TrimSpace(songName)
			songName = strings.Trim(songName, "-. ")
		}

		fmt.Printf("\t\ttrackNumber: %v\n", trackNumber)

		fmt.Printf("\t\tsong.Name(): %v\n", songName)
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

	err := findArtists(musicPath)
	if err != nil {
		log.Printf("error walking directories: %s", err.Error())
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

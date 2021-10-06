package service

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"

	"github.com/JoeRourke123/Monkey/handler"
	"github.com/JoeRourke123/Monkey/spec"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type Server struct {
	GRPCRoot     string
	HTTPRoot     string
	HTTPPort     string
	MusicPath    string
	DatabasePath string
	DB           *sqlx.DB
	StaticFS     embed.FS
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (s *Server) NotFound(w http.ResponseWriter, r *http.Request) {
	fileNotFoundBytes, err := fs.ReadFile(s.StaticFS, "dist/404.html")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(fileNotFoundBytes)
}

func (s *Server) FileServerWithCustom404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			s.NotFound(w, r)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

func (s *Server) NewAppMux() (*http.ServeMux, error) {
	distRoot, err := fs.Sub(s.StaticFS, "dist")
	if err != nil {
		log.Fatal(err)
	}

	appMux := http.NewServeMux()
	appMux.Handle("/", http.StripPrefix("/", s.FileServerWithCustom404(http.FS(distRoot))))

	return appMux, nil
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.GRPCRoot)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	h := &handler.Handler{
		DB:        s.DB,
		HTTPRoot:  s.HTTPRoot,
		MusicPath: s.MusicPath,
	}

	spec.RegisterMonkeyServer(grpcServer, h)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	apiMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = spec.RegisterMonkeyHandlerFromEndpoint(ctx, apiMux, s.GRPCRoot, opts)
	if err != nil {
		return err
	}

	s.customRoutes(apiMux, h)

	appMux, _ := s.NewAppMux()

	rootMux := http.NewServeMux()
	rootMux.Handle("/", appMux)
	rootMux.Handle("/v1/", apiMux)

	return http.ListenAndServe(s.HTTPPort, cors(rootMux))
}

func (s *Server) customRoutes(mux *runtime.ServeMux, h *handler.Handler) {
	err := mux.HandlePath("GET",
		"/v1/artist/{artist_id}/album/{album_id}/art",
		func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			var path string
			err := h.DB.Get(&path, `
				SELECT album.album_art FROM album
				WHERE album.id=$1 AND album.artist_id=$2;
				`, params["album_id"], params["artist_id"])
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusNotFound)
					log.Printf("illegal access '%s'", err.Error())
					return
				}

				w.WriteHeader(http.StatusBadRequest)
				log.Printf("error finding path '%s'", err.Error())
				return
			}

			absolutePath := s.MusicPath + path

			fileBytes, err := os.ReadFile(absolutePath)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}

			w.Write(fileBytes)
		})
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("GET",
		"/v1/artist/{artist_id}/album/{album_id}/song/{song_id}/media",
		func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			var path string
			err := h.DB.Get(&path, `
			SELECT song.path FROM song
			INNER JOIN album ON song.album_id=album.id
			WHERE song.id=$1 AND album.id=$2 AND album.artist_id=$3;
			`, params["song_id"], params["album_id"], params["artist_id"])
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusNotFound)
					log.Printf("illegal access '%s'", err.Error())
					return
				}

				w.WriteHeader(http.StatusBadRequest)
				log.Printf("error finding path '%s'", err.Error())
				return
			}

			absolutePath := s.MusicPath + path

			file, err := os.Open(absolutePath)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}

			info, err := file.Stat()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			http.ServeContent(w, r, absolutePath, info.ModTime(), file)
		})
	if err != nil {
		panic(err)
	}
}

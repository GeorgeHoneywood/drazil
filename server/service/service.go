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
	"time"

	"github.com/GeorgeHoneywood/drazil/handler"
	"github.com/GeorgeHoneywood/drazil/spec"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	GRPCRoot     string
	HTTPPort     string
	MusicPath    string
	DatabasePath string
	DB           *sqlx.DB
	StaticFS     embed.FS
	Log          *zap.Logger
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

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		s.Log.Debug("http request", zap.String("method", req.Method), zap.String("uri", req.RequestURI), zap.Duration("latency", time.Since(start)))
	})
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
		MusicPath: s.MusicPath,
		Log:       s.Log,
	}

	spec.RegisterDrazilServer(grpcServer, h)

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
	err = spec.RegisterDrazilHandlerFromEndpoint(ctx, apiMux, s.GRPCRoot, opts)
	if err != nil {
		return err
	}

	s.customRoutes(apiMux, h)

	appMux, _ := s.NewAppMux()

	rootMux := http.NewServeMux()
	rootMux.Handle("/", appMux)
	rootMux.Handle("/v1/", apiMux)

	loggingMiddleware := s.loggingMiddleware(rootMux)
	corsMiddleware := cors(loggingMiddleware)

	s.Log.Info("serving", zap.String("http", s.HTTPPort))
	return http.ListenAndServe(s.HTTPPort, corsMiddleware)
}

func (s *Server) customRoutes(mux *runtime.ServeMux, h *handler.Handler) {
	err := mux.HandlePath("GET",
		"/v1/artist/{artist_id}/album/{album_id}/art",
		func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			var artBytes []byte
			dbErr := h.DB.Get(&artBytes, `
				SELECT album.album_art FROM album
				WHERE album.id=$1 AND album.artist_id=$2;
				`, params["album_id"], params["artist_id"])
			if dbErr != nil {
				if dbErr == sql.ErrNoRows {
					w.WriteHeader(http.StatusNotFound)
					s.Log.Warn("art not found", zap.Error(dbErr), zap.String("album_id", params["album_id"]))
					return
				}

				w.WriteHeader(http.StatusBadRequest)
				s.Log.Error("error loading art", zap.Error(dbErr))
				return
			}

			if len(artBytes) == 0 {
				w.WriteHeader(http.StatusNotFound)
				// s.Log.Warn("art not found",
				// 	zap.String("album_id", params["album_id"]),
				// )
				return
			}

			w.Write(artBytes)
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

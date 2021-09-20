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

func (s *Server) ServeApp() error {
	distRoot, err := fs.Sub(s.StaticFS, "dist")
	if err != nil {
		log.Fatal(err)
	}

	appMux := http.NewServeMux()
	appMux.Handle("/", http.StripPrefix("/", s.FileServerWithCustom404(http.FS(distRoot))))

	appServer := &http.Server{
		Addr:         ":3001",
		Handler:      appMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return appServer.ListenAndServe()
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

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = spec.RegisterMonkeyHandlerFromEndpoint(ctx, mux, s.GRPCRoot, opts)
	if err != nil {
		return err
	}

	s.customRoutes(mux, h)

	go func() {
		err := s.ServeApp()
		fmt.Printf("err: %v\n", err)
	}()

	return http.ListenAndServe(s.HTTPPort, cors(mux))
}

func (s *Server) customRoutes(mux *runtime.ServeMux, h *handler.Handler) {
	// mux.HandlePath("GET", "/app/{path=**}", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
	// 	path := params["path"]

	// 	if path == "" {
	// 		path = "index.html"
	// 	}

	// 	fileBytes, err := fs.ReadFile(s.StaticFS, "dist/"+path) // FIXME: can escape here

	// 	if err != nil {
	// 		fmt.Printf("err: %v\n", err)
	// 		fileNotFoundBytes, err := fs.ReadFile(s.StaticFS, "dist/404.html")

	// 		if err != nil {
	// 			fmt.Printf("err: %v\n", err)
	// 			w.WriteHeader(http.StatusInternalServerError)
	// 			return
	// 		}

	// 		w.Write(fileNotFoundBytes)
	// 		return
	// 	}

	// 	w.Write(fileBytes)
	// })

	// err = mux.HandlePath("GET", "/media/{path=**}", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
	// 	absolutePath := s.MusicPath + params["path"]

	// 	var path string
	// 	err := h.DB.Get(&path, `
	// 	SELECT path FROM
	// 	(
	// 		SELECT album_art AS path FROM album
	// 		UNION
	// 		SELECT path FROM song
	// 	)
	// 	WHERE path=$1;
	// 	`, params["path"])
	// 	if err != nil {
	// 		if err == sql.ErrNoRows {
	// 			w.WriteHeader(http.StatusNotFound)
	// 			log.Printf("illegal access '%s': %s", absolutePath, err.Error())
	// 			return
	// 		}

	// 		w.WriteHeader(http.StatusBadRequest)
	// 		log.Printf("error finding path '%s': %s", absolutePath, err.Error())
	// 		return
	// 	}

	// 	fileBytes, err := os.ReadFile(absolutePath)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusNotFound)
	// 	}

	// 	w.Write(fileBytes)
	// })
	// if err != nil {
	// 	panic(err)
	// }

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

			fileBytes, err := os.ReadFile(absolutePath)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}

			w.Write(fileBytes)
		})
	if err != nil {
		panic(err)
	}
}

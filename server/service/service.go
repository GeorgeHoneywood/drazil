package service

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"

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

	err = mux.HandlePath("GET", "/media/{path=**}", func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		absolutePath := s.MusicPath + params["path"]

		var path string
		err := h.DB.Get(&path, `
		SELECT path FROM
		(
			SELECT album_art AS path FROM album
			UNION
			SELECT path FROM song
		)
		WHERE path=$1;
		`, params["path"])
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				log.Printf("illegal access '%s': %s", absolutePath, err.Error())
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			log.Printf("error finding path '%s': %s", absolutePath, err.Error())
			return
		}

		fileBytes, err := os.ReadFile(absolutePath)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Write(fileBytes)
	})
	if err != nil {
		panic(err)
	}

	return http.ListenAndServe(s.HTTPPort, cors(mux))
}

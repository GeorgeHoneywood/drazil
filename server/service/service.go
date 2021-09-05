package service

import (
	"context"
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
		fileBytes, err := os.ReadFile(s.MusicPath + params["path"])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Write(fileBytes)
	})
	if err != nil {
		panic(err)
	}

	return http.ListenAndServe(s.HTTPRoot, cors(mux))
}

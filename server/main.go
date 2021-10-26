package main

import (
	"embed"
	"flag"
	"fmt"

	"github.com/JoeRourke123/Monkey/models"
	"github.com/JoeRourke123/Monkey/scanner"
	"github.com/JoeRourke123/Monkey/service"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	_ "modernc.org/sqlite"
)

//go:embed dist/* migrations/*.sql

var static embed.FS

func main() {
	musicPath := flag.String("d", "/home/honeyfox/Music/", "music path (include trailing slash)")
	flag.Parse()

	log, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("error creating logger")
	}

	s := &service.Server{
		GRPCRoot:     "localhost:9091",
		HTTPPort:     "0.0.0.0:4444",
		MusicPath:    *musicPath,
		DatabasePath: "./monkey.db",
		StaticFS:     static,
		Log:          log,
	}

	// TESTING ONLY
	// os.Remove(s.DatabasePath)

	err = models.SetupTables(s.DatabasePath, s.StaticFS)
	if err != nil {
		s.Log.Debug("error setting up db", zap.Error(err))
	}

	s.Log.Info("connecting to database", zap.String("path", s.DatabasePath))
	s.DB, err = sqlx.Connect("sqlite", s.DatabasePath)
	if err != nil {
		s.Log.Fatal("could not connect to database", zap.Error(err))
	}
	defer func() {
		err := s.DB.Close()
		if err != nil {
			s.Log.Fatal("could not close database connection", zap.Error(err))
		}
	}()

	s.Log.Debug("scanning for music", zap.String("path", s.MusicPath))
	sc := &scanner.Scanner{
		MusicPath: s.MusicPath,
		Log:       s.Log,
	}

	go func() {
		err = sc.FindArtists(s.DB)
		if err != nil {
			s.Log.Debug("error scanning music", zap.Error(err))
		}
	}()

	if err := s.Run(); err != nil {
		s.Log.Fatal("fatal error during run", zap.Error(err))
	}
}

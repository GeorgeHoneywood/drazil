package main

import (
	"embed"
	"fmt"

	"github.com/JoeRourke123/Monkey/models"
	"github.com/JoeRourke123/Monkey/scanner"
	"github.com/JoeRourke123/Monkey/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

//go:embed dist/* migrations/*.sql

var static embed.FS

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("error creating logger")
	}

	s := &service.Server{
		GRPCRoot:     "localhost:9091",
		HTTPPort:     "0.0.0.0:4444",
		MusicPath:    "/home/honeyfox/Music/",
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
	s.DB, err = sqlx.Connect("sqlite3", s.DatabasePath)
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

	err = sc.FindArtists(s.DB)
	if err != nil {
		s.Log.Debug("error scanning music", zap.Error(err))
	}

	if err := s.Run(); err != nil {
		s.Log.Fatal("fatal error during run", zap.Error(err))
	}
}

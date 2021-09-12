package main

import (
	"errors"
	"log"
	"os"

	"github.com/JoeRourke123/Monkey/models"
	"github.com/JoeRourke123/Monkey/scanner"
	"github.com/JoeRourke123/Monkey/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := &service.Server{
		GRPCRoot:     "localhost:9091",
		HTTPPort:     "localhost:8081",
		HTTPRoot:     "http://localhost:8081",
		MusicPath:    "/home/honeyfox/Music/",
		DatabasePath: "./monkey.db",
	}

	os.Remove(s.DatabasePath)

	_, err := os.Stat(s.DatabasePath)
	exist := !errors.Is(err, os.ErrNotExist)

	s.DB, err = sqlx.Connect("sqlite3", s.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := s.DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if !exist {
		log.Print("setting up tables")

		err = models.SetupTables(s.DB)
		if err != nil {
			log.Fatalf("error setting up db: %s", err.Error())
		}
	}

	err = scanner.FindArtists(s.MusicPath, s.DB)
	if err != nil {
		log.Printf("error walking directories: %s", err.Error())
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}

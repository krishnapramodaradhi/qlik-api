package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/krishnapramodaradhi/qlik-api/internal/config"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("There was an error while loading the env variables, err")
	}
}

func main() {
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatal("there was an error while connecting to the db", err)
	}
	s := config.NewServer(":8080", db.DB)
	s.Run()
}

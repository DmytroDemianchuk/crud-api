package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dmytrodemianchuk/crud-api/internal/repository"
	"github.com/dmytrodemianchuk/crud-api/internal/service"
	"github.com/dmytrodemianchuk/crud-api/internal/transport/rest"
	"github.com/dmytrodemianchuk/crud-api/pkg/database"

	_ "github.com/lib/pq"
)

func main() {
	// init db
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "goLANGn1nja",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// // init deps
	musicsRepo := repository.NewMusic(db)
	musicService := service.NewMusic(musicsRepo)
	handler := rest.NewMusic(musicService)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"birthdays/internal/auth"
	"birthdays/internal/config"
	"birthdays/internal/handlers"
	"birthdays/internal/logging"
	"birthdays/internal/mux"
	"birthdays/internal/services"
	"birthdays/internal/storages"
	postgresql "birthdays/internal/storages/postgres"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

const envPath = "config/.env"
const configPath = "config/config.yml"

func main() {
	config, err := config.LoadConfig(configPath, envPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(time.Now(), "Config loaded")

	loggers, err := logging.Setup(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(time.Now(), "Logger set up")

	db, err := postgresql.GetDBConnection(*config.Database)
	if err != nil {
		log.Fatal(err.Error())
	}

	authManager := auth.NewManager(config.JWT)

	storages := storages.NewPostgresStorages(db)
	services := services.NewServices(storages, authManager)
	handlers := handlers.NewHandlers(services)

	mux, err := mux.SetUpMux(handlers, loggers.HTTP)
	if err != nil {
		log.Fatal(err.Error())
	}

	var server = http.Server{
		Addr:    fmt.Sprintf(":%d", config.App.Port),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	server.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

package main

import (
	"birthdays/internal/auth"
	"birthdays/internal/config"
	"birthdays/internal/handlers"
	"birthdays/internal/logging"
	"birthdays/internal/mux"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/services"
	"birthdays/internal/storages"
	postgresql "birthdays/internal/storages/postgres"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-co-op/gocron/v2"

	"log"
)

const envPath = "config/.env"
const configPath = "config/config.yml"

// @title           Rutube Birthdays
// @version         1.0
// @description     Тестовое задание: Бэкенд-сервер для учёта дней рождения

// @contact.name   Никита Архаров

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  JWT
// @in header
// @name "Authorization"
func main() {
	config, err := config.LoadConfig(configPath, envPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(time.Now(), "Config loaded")

	logger, err := logging.Setup(config.Logging)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(time.Now(), "Logger set up")

	db, err := postgresql.GetDBConnection(*config.Database)
	if err != nil {
		log.Fatal(err.Error())
	}

	authManager := auth.NewManager(config.JWT)
	scheduler, _ := gocron.NewScheduler()

	storages := storages.NewPostgresStorages(db)
	services := services.NewServices(storages, authManager, scheduler)
	handlers := handlers.NewHandlers(services)

	jobContext := context.Background()
	go func(ctx context.Context) {
		ctx = context.WithValue(ctx, dto.CtxLoggerKey, logger.Logger)
		err := services.Job.Gather(ctx)
		if err != nil {
			log.Print(err.Error())
		}
		services.Job.Start()
	}(jobContext)

	mux, err := mux.SetUpMux(handlers, logger)
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

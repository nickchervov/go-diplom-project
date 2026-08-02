package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nickchervov/go-diplom-project/internal/adapter/sqlite"
	"github.com/nickchervov/go-diplom-project/internal/controller/handler"
	"github.com/nickchervov/go-diplom-project/internal/service"
	"github.com/nickchervov/go-diplom-project/pkg/httpserver"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("loading config: %v", err)
	}

	if err := AppRun(context.Background()); err != nil {
		log.Fatalf("running app: %v", err)
	}
}

func AppRun(ctx context.Context) error {
	sqlite, err := sqlite.New(ctx)
	if err != nil {
		return fmt.Errorf("open connection to db: %w", err)
	}

	svc := service.New(sqlite)

	router := handler.SetRoutes(svc)

	server := httpserver.New(router)

	go func() {
		log.Println("Starting server on port:", server.Server.Addr)
		if err := server.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("connection to server: %v", err)
		}
	}()

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
	log.Println("Detected graceful shutdown signal")
	log.Println("Starting Graceful shutdown")

	server.Close()
	sqlite.Close()

	log.Println("Server closed")
	return nil
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/akselarzuman/go-jaeger/api/router"
	"github.com/akselarzuman/go-jaeger/internal/pkg/db"
	"github.com/akselarzuman/go-jaeger/internal/pkg/opentelemetry"
	"github.com/akselarzuman/go-jaeger/internal/pkg/redis"
	"github.com/joho/godotenv"
)

func main() {
	initEnv()

	tp, err := opentelemetry.NewJaegerTraceProvider()
	if err != nil {
		log.Println(err.Error())
	}

	if tp != nil {
		defer func() {
			// Cleanly shutdown and flush telemetry when the application exits.
			exitCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := tp.Shutdown(exitCtx); err != nil {
				log.Println(err.Error())
			}
		}()
	}

	initResources()

	app := router.Setup()

	// graceful shutdown
	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()

	<-signalCh

	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Server stopped gracefully...")
}

func initEnv() {
	if err := godotenv.Load(path.Join(getRootPath(), "/.env")); err != nil {
		log.Println("Error while opening .env file", err.Error())
	}
}

func initResources() {
	if p := db.NewPostgresClient(); p != nil {
		log.Println("postgres client created")

		if err := p.AutoMigrate(); err != nil {
			log.Println("error while auto migrating", err.Error())
		}
	}

	if r := redis.NewRedisConnection(); r != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := r.Ping(ctx).Err(); err != nil {
			log.Println("error while pinging redis", err.Error())
		} else {
			log.Println("redis connection created")
		}
	} else {
		log.Println("redis connection failed")
	}
}

func getRootPath() string {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(p, "../")
}

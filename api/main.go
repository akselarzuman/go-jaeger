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
	"github.com/akselarzuman/go-jaeger/internal/pkg/jaeger"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
)

func main() {
	initEnv()
	tracer, closer := jaeger.Init("uber-api")
	defer closer.Close()
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)

	app := router.Setup()

	// graceful shutdown
	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":1000",
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

func getRootPath() string {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(p, "../")
}

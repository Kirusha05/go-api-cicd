package main

import (
	"context"
	"log"
	"my-api/internal/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	address string = ":80"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	// Routes
	mux.Handle("/users/", http.StripPrefix("/users", api.NewMux()))

	// Start the server
	log.Default().Printf("Server listening on port %v", address)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error: ", err)
		}
	}()

	// Setup a channel to listen for SIGINT or SIGTERM
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	signal := <-stopSignal // Wait for a signal and store the value
	log.Default().Println("Received signal:", signal.String())

	// Graceful shutdown
	log.Default().Println("Shutting down server...")
	startTime := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed: ", err)
	}
	log.Default().Println("Server stopped gracefully in", time.Since(startTime))
}

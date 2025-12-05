package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"awesomeProject/internal/handlers"
	"awesomeProject/internal/services"
	"awesomeProject/internal/store"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Create logger
	logger := log.Default() // using standard log; replace with your custom log if needed

	// Create TCP store (adjust address to your TCP server)
	tcpStore := store.NewTCPStore("127.0.0.1:6379")

	// Create KV service
	kvSvc := services.NewKVService(tcpStore, logger)

	// Put services into struct
	svcs := &services.Services{
		KV: kvSvc,
	}

	// Create router
	router := chi.NewRouter()

	// Register KV routes
	handlers.RegisterKVRoutes(router, svcs.KV)

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Run server in goroutine
	go func() {
		log.Println("HTTP server listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server stopped gracefully")
}

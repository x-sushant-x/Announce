package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	registry := NewRegistry()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/service", makeHandler(registry))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	errChan := make(chan error, 1)

	go func() {
		log.Println("🚀 Server is running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Printf("❌ Server error: %v", err)

	case sig := <-quitChan:
		log.Printf("🛑 Received OS signal: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("🔄 Shutting down server gracefully...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("⚠️ Could not gracefully shutdown: %v", err)
	}
	log.Println("✅ Server shutdown complete.")
}

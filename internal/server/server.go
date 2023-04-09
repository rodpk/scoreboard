package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateServer() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	return router
}

func StartServer(r *chi.Mux) {

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	go func() {
		log.Println("Starting server on port 8000...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// create a channel to receive OS signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// block until we receive a signal
	<-sig

	log.Println("Shutting down server...")

	// create a context with timeout to force shutdown if it takes too long
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	//shutdown server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped.")
}

func CloseServer() {

}

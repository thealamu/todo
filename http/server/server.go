package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/sethvargo/go-signalcontext"
)

// Server wraps a http server providing more methods
type Server struct {
	srv *http.Server
}

// New returns a newly configured http server
func New(addr string, router http.Handler) *Server {
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return &Server{srv}
}

// Shutdown gracefully shuts down the http server
func (s Server) Shutdown() {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}

// Run starts the http server
func (s Server) Run() {
	ctx, cancel := signalcontext.OnInterrupt()
	defer cancel()

	log.Println("Starting server on", s.srv.Addr)
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	s.Shutdown()
}

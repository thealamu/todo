package server

import (
	"net/http"
	"time"
)

// New returns a newly configured http server
func New(addr string, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}

package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// We read port from the environment but default to 1028
	port := os.Getenv("PORT")
	if port == "" {
		port = "1028"
	}
	// Build the address to run on
	srvAddr := net.JoinHostPort("", port)
	srv := &http.Server{
		Addr:         srvAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

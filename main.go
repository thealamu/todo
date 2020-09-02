package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/thealamu/todo/create"
	"github.com/thealamu/todo/delete"
	"github.com/thealamu/todo/retrieve"
	"github.com/thealamu/todo/update"
)

var logger = log.New(os.Stdout, "To-Do API:", log.LstdFlags|log.Lshortfile)

func main() {
	// We read port from the environment but default to 1028
	port := os.Getenv("PORT")
	if port == "" {
		port = "1028"
	}

	// Build the router
	router := mux.NewRouter()
	// Delegate route registering whilst injecting a logger
	create.New(logger).Register(router)
	retrieve.New(logger).Register(router)
	update.New(logger).Register(router)
	delete.New(logger).Register(router)

	// Build the address to run on
	srvAddr := net.JoinHostPort("", port)
	srv := &http.Server{
		Addr:         srvAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

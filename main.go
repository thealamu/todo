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

func main() {
	// We read port from the environment but default to 1028
	port := os.Getenv("PORT")
	if port == "" {
		port = "1028"
	}

	// Build the router
	router := mux.NewRouter()
	// Delegate route registering
	create.New().Register(router)
	retrieve.New().Register(router)
	update.New().Register(router)
	delete.New().Register(router)

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

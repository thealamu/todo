package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/thealamu/todo/create"
	"github.com/thealamu/todo/db"
	"github.com/thealamu/todo/delete"
	"github.com/thealamu/todo/retrieve"
	"github.com/thealamu/todo/update"
)

var logger = log.New(os.Stdout, "To-Do API: ", log.LstdFlags|log.Lshortfile)

func main() {
	// Build the router
	router := mux.NewRouter()
	// Delegate route registering and inject dependencies
	var memDB = db.NewInMem()
	create.New(logger, memDB).Register(router)
	retrieve.New(logger, memDB).Register(router)
	update.New(logger, memDB).Register(router)
	delete.New(logger, memDB).Register(router)

	srvAddr := getSrvAddress()
	srv := &http.Server{
		Addr:         srvAddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Start the server
	logger.Println("Starting the server on", srvAddr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func getSrvAddress() string {
	// We read port from the environment but default to 1028
	port := os.Getenv("PORT")
	if port == "" {
		port = "1028"
	}
	return net.JoinHostPort("", port)
}

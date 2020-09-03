package main

import (
	"log"
	"net"
	"os"

	"github.com/gorilla/mux"
	"github.com/thealamu/todo/create"
	"github.com/thealamu/todo/db"
	"github.com/thealamu/todo/delete"
	"github.com/thealamu/todo/http/server"
	"github.com/thealamu/todo/retrieve"
	"github.com/thealamu/todo/update"
)

var logger = log.New(os.Stdout, "To-Do API: ", log.LstdFlags|log.Lshortfile)
var memDB = db.NewInMem()

func main() {
	// Build the router
	router := mux.NewRouter()
	// Delegate route registering and inject dependencies
	create.New(logger, memDB).Register(router)
	retrieve.New(logger, memDB).Register(router)
	update.New(logger, memDB).Register(router)
	delete.New(logger, memDB).Register(router)

	srv := server.New(getSrvAddress(), router)
	srv.Run()
}

func getSrvAddress() string {
	// We read port from the environment but default to 1028
	port := os.Getenv("PORT")
	if port == "" {
		port = "1028"
	}
	return net.JoinHostPort("", port)
}

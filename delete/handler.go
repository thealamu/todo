package delete

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thealamu/todo/db"
)

// Handler handles all item creation
type Handler struct {
	logger *log.Logger
	db     db.DB
}

// New returns a new Handler
func New(logger *log.Logger, db db.DB) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

// Register registers the endpoints in mux
func (h *Handler) Register(mux *mux.Router) {
	mux.HandleFunc("/todos/{id:[0-9]+}", h.DeleteSingle).Methods(http.MethodDelete)
}

// DeleteSingle deletes a single to-do item
func (h *Handler) DeleteSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Deleting Single")
}

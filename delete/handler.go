package delete

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler handles all item creation
type Handler struct {
	logger *log.Logger
}

// New returns a new Handler
func New(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// Register registers the endpoints in mux
func (h *Handler) Register(mux *mux.Router) {
	mux.HandleFunc("/todos/{id}", h.DeleteSingle).Methods(http.MethodDelete)
}

// DeleteSingle deletes a single to-do item
func (h *Handler) DeleteSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Deleting Single")
}

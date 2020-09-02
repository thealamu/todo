package retrieve

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
	mux.HandleFunc("/todos", h.GetAll).Methods(http.MethodGet)
	mux.HandleFunc("/todos/{id:[0-9]+}", h.GetSingle).Methods(http.MethodGet)
}

// GetAll serves all to-do items
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Getting All")
}

// GetSingle serves a single to-do item
func (h *Handler) GetSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Getting Single")
}

package update

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler handles all item creation
type Handler struct{}

// New returns a new Handler
func New() *Handler {
	return &Handler{}
}

// Register registers the endpoints in mux
func (h *Handler) Register(mux *mux.Router) {
	mux.HandleFunc("/todos/{id}", h.UpdateSingle).Methods(http.MethodPut)
}

// UpdateSingle updates a single to-do item
func (h *Handler) UpdateSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Updating Single")
}

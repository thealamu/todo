package update

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thealamu/todo/db"
	"github.com/thealamu/todo/http/respond"
	"github.com/thealamu/todo/todo"
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
	mux.HandleFunc("/todos/{id:[0-9]+}", h.UpdateSingle).Methods(http.MethodPut)
}

// UpdateSingle updates a single to-do item
func (h *Handler) UpdateSingle(w http.ResponseWriter, r *http.Request) {
	var td todo.Todo
	err := json.NewDecoder(r.Body).Decode(&td)
	if err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	// Update item
	h.logger.Println("Updating to-do item with ID", td.ID)
	err = h.db.UpdateItem(td)
	if err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
}

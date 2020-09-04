package create

import (
	"encoding/json"
	"fmt"
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
	mux.HandleFunc("/todos", h.CreateSingle).Methods(http.MethodPost)
}

// CreateSingle creates a single to-do item
func (h *Handler) CreateSingle(w http.ResponseWriter, r *http.Request) {
	var item todo.Todo
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	// Set the next ID on the to-do item
	item.ID = h.db.GetNextID()
	// Save the to-do item
	h.logger.Println("Creating to-do item with ID", item.ID)
	h.db.AddItem(item)
	// Return Location of item
	itemLoc := r.Host + fmt.Sprintf("/todos/%d", item.ID)
	w.Header().Set("Location", itemLoc)
	// Return full to-do item
	respond.JSON(w, item, http.StatusCreated)
}

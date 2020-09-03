package retrieve

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thealamu/todo/db"
	"github.com/thealamu/todo/http/respond"
)

// Handler handles all item retrieval
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
	mux.HandleFunc("/todos", h.GetAll).Methods(http.MethodGet)
	mux.HandleFunc("/todos/{id:[0-9]+}", h.GetSingle).Methods(http.MethodGet)
}

// GetAll serves all to-do items
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Getting all to-do items")
	items := h.db.GetAllItems()
	h.logger.Printf("Returning %d to-do items\n", len(items))
	// Write items to output as JSON
	err := respond.JSON(w, items)
	if err != nil {
		log.Fatal(err)
	}
}

// GetSingle serves a single to-do item
func (h *Handler) GetSingle(w http.ResponseWriter, r *http.Request) {
	// Read id variable from path
	varID := mux.Vars(r)["id"]
	todoID, err := strconv.Atoi(varID)
	if err != nil {
		log.Fatal(err)
	}

	h.logger.Printf("Getting single to-do item for ID %d", todoID)
	todoItem, err := h.db.GetSingleItem(todoID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = respond.JSON(w, todoItem)
	if err != nil {
		log.Fatal(err)
	}
}

package retrieve

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thealamu/todo/db"
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
	// Write JSON of items to output
	err := h.RespondJSON(w, items)
	if err != nil {
		log.Fatal(err)
	}
}

// GetSingle serves a single to-do item
func (h *Handler) GetSingle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Getting Single")
}

// RespondJSON responds with data as JSON
func (h *Handler) RespondJSON(w http.ResponseWriter, data ...interface{}) error {
	coded, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Write(coded)
	return nil
}

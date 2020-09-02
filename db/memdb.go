package db

import "github.com/thealamu/todo/todo"

// MemoryDB is an in-memory implementation of the DB interface
type MemoryDB struct{}

var inMemTodos []todo.Todo

// NewInMem returns an in-memory DB
func NewInMem() *MemoryDB {
	return &MemoryDB{}
}

// GetAllItems returns all to-do items
func (m *MemoryDB) GetAllItems() []todo.Todo {
	return inMemTodos
}

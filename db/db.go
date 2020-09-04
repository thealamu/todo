package db

import "github.com/thealamu/todo/todo"

// DB interafce holds common methods for db implementations
type DB interface {
	GetAllItems() []todo.Todo
	GetSingleItem(id int) (todo.Todo, error)
	GetNextID() int
	AddItem(todo.Todo)
	DeleteItem(id int) error
}

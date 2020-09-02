package db

import "github.com/thealamu/todo/todo"

// DB interafce holds common methods for db implementations
type DB interface {
	GetAllItems() []todo.Todo
}

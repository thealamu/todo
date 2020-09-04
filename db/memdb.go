package db

import (
	"fmt"
	"sync"

	"github.com/thealamu/todo/todo"
)

// MemoryDB is an in-memory implementation of the DB interface
type MemoryDB struct{}

var l sync.Mutex
var inMemTodos []todo.Todo

// NewInMem returns an in-memory DB
func NewInMem() *MemoryDB {
	return &MemoryDB{}
}

// GetAllItems returns all to-do items
func (m *MemoryDB) GetAllItems() []todo.Todo {
	return inMemTodos
}

// GetSingleItem returns a to-do item with the given id
func (m *MemoryDB) GetSingleItem(id int) (todo.Todo, error) {
	i, err := m.findIndexForID(id)
	if err != nil {
		return todo.Todo{}, err
	}
	return inMemTodos[i], nil
}

// GetNextID returns the next to-do ID
func (m *MemoryDB) GetNextID() int {
	hIndx := len(inMemTodos) - 1
	if hIndx < 0 {
		return 1
	}
	return inMemTodos[hIndx].ID + 1
}

// AddItem adds a to-do item to the list
func (m *MemoryDB) AddItem(td todo.Todo) {
	l.Lock()
	defer l.Unlock()
	inMemTodos = append(inMemTodos, td)
}

// DeleteItem removes to-do item with id from the list
func (m *MemoryDB) DeleteItem(id int) error {
	l.Lock()
	defer l.Unlock()
	indx, err := m.findIndexForID(id)
	if err != nil {
		return err
	}
	var newInMem []todo.Todo
	newInMem = append(newInMem, inMemTodos[:indx]...)
	newInMem = append(newInMem, inMemTodos[indx+1:]...)
	inMemTodos = newInMem
	return nil
}

// FindIndexForID returns the index of a to-do item, given the ID.
// It searches the slice of to-dos using the binary search algorithm.
// The binary search algorithm is an efficient searching algorithm with a
// worst case O(logn), the only requisite is the items must be sorted.
// This is not a problem because new IDs are generated by incrementing
// the last ID by 1. So naturally, adding or removing items maintains the invariance
// which is, at any point in time, the inMemTodos slice is sorted.
func (m *MemoryDB) findIndexForID(id int) (int, error) {
	i, j := 0, len(inMemTodos)-1
	for i <= j {
		m := (i + j) / 2
		if inMemTodos[m].ID < id {
			i = m + 1
		} else if inMemTodos[m].ID > id {
			j = m - 1
		} else {
			return m, nil
		}
	}
	return -1, fmt.Errorf("Not Found")
}

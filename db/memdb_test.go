package db

import (
	"testing"

	"github.com/thealamu/todo/todo"
)

func TestDeleteItem(t *testing.T) {
	inMemTodos = append(inMemTodos, todo.Todo{ID: 1}, todo.Todo{ID: 2}, todo.Todo{ID: 3})
	m := NewInMem()
	err := m.DeleteItem(2)
	if err != nil {
		t.Error(err)
	}
	_, err = m.GetSingleItem(2)
	if err == nil {
		t.Errorf("DeleteItem does not remove to-do item, we can still retrieve it")
	}

	err = m.DeleteItem(0)
	if err == nil {
		t.Errorf("DeleteItem returns nil error for non-existent to-do item")
	}
}

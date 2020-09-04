package db

import (
	"testing"

	"github.com/thealamu/todo/todo"
)

func TestUpdateItem(t *testing.T) {
	inMemTodos = append(inMemTodos, todo.Todo{ID: 1, Done: false})
	m := NewInMem()
	upted := todo.Todo{ID: 1, Done: true}
	err := m.UpdateItem(upted)
	if err != nil {
		t.Error(err)
	}
	if inMemTodos[0].Done != upted.Done {
		t.Errorf("UpdateItem does not do item update")
	}
	err = m.UpdateItem(todo.Todo{ID: 3})
	if err == nil {
		t.Errorf("UpdateItem returns nil error for non-existent to-do item")
	}
}

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

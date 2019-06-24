package main

import (
	"github.com/go-pg/pg/orm"
)

var todos Todos

// FindTodo finds a todo
func FindTodo(id uint64) (*Todo, error) {
	t := &Todo{
		ID: id,
	}
	err := db.Select(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

//CreateTodo creates a Todo
func CreateTodo(t *Todo) (*Todo, error) {
	err := db.Insert(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// RepoDestroyTodo destroys a todo
func RepoDestroyTodo(id uint64) error {
	t := &Todo{
		ID: id,
	}
	return db.Delete(t)
}

func createSchema() error {
	model := (*Todo)(nil)
	err := db.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}

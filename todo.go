package main

import "time"

// Todo struct
type Todo struct {
	ID        uint64    `json:"id" sql:",notnull"`
	Name      string    `json:"name" sql:",notnull"`
	Completed bool      `json:"completed" sql:"default:false,notnull"`
	Due       time.Time `json:"due"`
}

// Todos is a list of Todo
type Todos []Todo

package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	addTodo(name string) *Todo
	// Close()
}

var handler dbHandler

func init() {
	handler = newMemoryHandler()
}

func AddTodo(name string) *Todo {
	return handler.addTodo(name)
}

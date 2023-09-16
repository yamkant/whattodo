package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	getTodos() []*Todo
	addTodo(content string) *Todo
	removeTodo(id int) bool
	// Close()
}

var handler dbHandler

func init() {
	handler = newMemoryHandler()
}

func GetTodos() []*Todo {
	return handler.getTodos()
}

func AddTodo(content string) *Todo {
	return handler.addTodo(content)
}

func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}

package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	getTodos() []*Todo
	addTodo(content string) *Todo
	updateTodo(id int, completed bool, startedAt time.Time, endedAt time.Time) bool
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

func UpdateTodo(id int, completed bool, startedAt time.Time, endedAt time.Time) bool {
	return handler.updateTodo(id, completed, startedAt, endedAt)
}

func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}

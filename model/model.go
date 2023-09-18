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

type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(content string) *Todo
	UpdateTodo(id int, completed bool, startedAt time.Time, endedAt time.Time) bool
	RemoveTodo(id int) bool
	// Close()
}

func NewDBHandler() DBHandler {
	return newMemoryHandler()
}

package models

import "time"

type Todo struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Content   string    `json:"content"`
	Completed bool      `json:"completed"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoAddDTO struct {
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

type TodoUpdateDTO struct {
	Completed bool      `json:"completed"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
}

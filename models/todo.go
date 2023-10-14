package models

import "time"

type Todo struct {
	ID        int       `json:"id" gorm:"primary_key;column:id"`
	Content   string    `json:"content" gorm:"column:content"`
	Completed bool      `json:"completed" gorm:"column:completed"`
	StartedAt time.Time `json:"started_at" gorm:"column:started_at"`
	EndedAt   time.Time `json:"ended_at" gorm:"column:ended_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`

	UserID int
	User   User `gorm:"foreignKey:UserID"`
}

func (Todo) TableName() string {
	return "todos"
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

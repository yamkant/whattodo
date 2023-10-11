package model

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos(sessionId string) []*Todo {
	return nil
}

func (s *sqliteHandler) AddTodo(sessionId string, content string) *Todo {
	return nil
}

func (s *sqliteHandler) UpdateTodo(id int, completed bool, startedAt time.Time, endedAt time.Time) bool {
	return false
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	return false
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

// API USER
func (s *sqliteHandler) AddUser(email string, sessionId string) *User {
	return nil
}

func (s *sqliteHandler) GetUserBySessionId(sessionId string) *User {
	return nil
}

func newSqliteHandler() DBHandler {
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id			INTEGER  PRIMARY KEY AUTOINCREMENT,
			session_id	TEXT,
			content		BOOLEAN,
			completed	BOOLEAN,
			started_at	DATETIME,
			ended_at	DATETIME,
			created_at	DATETIME,
		)`)
	statement.Exec()
	return &sqliteHandler{db: database}
}

package model

import (
	"sort"
	"time"
)

type memoryHandler struct {
	todoMap map[int]*Todo
	userMap map[int]*User
}

// API TODO
func (m *memoryHandler) GetTodos(sessionId string) []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		if v.SessionId == sessionId {
			list = append(list, v)
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID > list[j].ID
	})
	return list
}

func (m *memoryHandler) AddTodo(sessionId string, content string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, sessionId, content, false, time.Time{}, time.Time{}, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) UpdateTodo(id int, completed bool, startedAt time.Time, endedAt time.Time) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.StartedAt = startedAt
		todo.EndedAt = endedAt
		todo.Completed = completed
		return true
	}
	return false
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

// API USER
func (m *memoryHandler) AddUser(email string, sessionId string) *User {
	id := len(m.userMap) + 1
	user := &User{id, email, sessionId, time.Now()}
	m.userMap[id] = user
	return user
}

func (m *memoryHandler) GetUserBySessionId(sessionId string) *User {
	var retUser *User
	for _, user := range m.userMap {
		if user.SessionId == sessionId {
			retUser = user
		}
	}
	return retUser
}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	m.userMap = make(map[int]*User)
	return m
}

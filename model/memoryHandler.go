package model

import (
	"fmt"
	"sort"
	"time"
)

type memoryHandler struct {
	todoMap map[int]*Todo
}

func (m *memoryHandler) GetTodos() []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		fmt.Print(v.Completed, v.StartedAt, v.EndedAt)
		list = append(list, v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID > list[j].ID
	})
	return list
}

func (m *memoryHandler) AddTodo(content string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, content, false, time.Time{}, time.Time{}, time.Now()}
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

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}

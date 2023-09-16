package model

import (
	"fmt"
	"sort"
	"time"
)

type memoryHandler struct {
	todoMap map[int]*Todo
}

func (m *memoryHandler) getTodos() []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		fmt.Println(v)
		list = append(list, v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID > list[j].ID
	})
	return list
}

func (m *memoryHandler) addTodo(content string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, content, false, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) removeTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func newMemoryHandler() dbHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}

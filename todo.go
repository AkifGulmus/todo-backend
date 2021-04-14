package main

import (
	"sort"
	"time"

	"github.com/google/uuid"
)

func GetTodos() []Todo {
	var todosList = []Todo{}
	for _, value := range Todos {
		todosList = append(todosList, Todo{createdDate: value.createdDate, TodoID: value.TodoID, Text: value.Text, Done: value.Done})
	}
	sort.Sort(byCreatedDate(todosList))
	return todosList
}

func CreateTodo(todo_text string) {
	todoID := uuid.New().String()
	Todos[todoID] = Todo{
		createdDate: time.Now().Unix(),
		TodoID:      todoID,
		Text:        todo_text,
		Done:        false,
	}
}

func UpdateTodo(todoID string) {
	todoToUpdate := Todos[todoID]
	todoToUpdate.Done = !todoToUpdate.Done
	Todos[todoID] = todoToUpdate
}

func DeleteTodo(todoID string) {
	delete(Todos, todoID)
}

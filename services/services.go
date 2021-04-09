package services

import (
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	TodoID string `json:"todoID"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
}

var Todos = make(map[string]Todo)

func GetTodos() []Todo {
	var todosList = []Todo{}
	for _, value := range Todos {
		todoItem := Todo{TodoID: value.TodoID, Text: value.Text, Done: value.Done}
		todosList = append(todosList, todoItem)
	}
	return todosList
}

func CreateTodo(todo_text string) {
	fmt.Println("NEVAR NE")
	todoID := uuid.New().String()
	Todos[todoID] = Todo{
		TodoID: todoID,
		Text:   todo_text,
		Done:   false,
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

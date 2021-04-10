package services

import (
	"sort"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	createdDate int64
	TodoID      string `json:"todoID"`
	Text        string `json:"text"`
	Done        bool   `json:"done"`
}

var Todos = make(map[string]Todo)

type byCreatedDate []Todo

func (a byCreatedDate) Len() int           { return len(a) }
func (a byCreatedDate) Less(i, j int) bool { return a[i].createdDate < a[j].createdDate }
func (a byCreatedDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

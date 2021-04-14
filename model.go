package main

type Todo struct {
	createdDate int64
	TodoID      string `json:"todoID"`
	Text        string `json:"text"`
	Done        bool   `json:"done"`
}

var Todos = make(map[string]Todo)

type requestBodyModel struct {
	Text string
	Done bool
}

type byCreatedDate []Todo

func (a byCreatedDate) Len() int           { return len(a) }
func (a byCreatedDate) Less(i, j int) bool { return a[i].createdDate < a[j].createdDate }
func (a byCreatedDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodo struct {
	Text string `json:"text"`
}

type SelectTodo struct {
	TodoID string `json:"todoId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

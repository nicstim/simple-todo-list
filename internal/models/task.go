package models

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

package handler

import (
	"fmt"
	"net/http"
)

type Todo struct{}

func (t *Todo) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Todo")
}

func (t *Todo) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all todos")
}

func (t *Todo) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Todo By ID")
}

func (t *Todo) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Todo By ID")
}

func (t *Todo) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Todo By ID")
}

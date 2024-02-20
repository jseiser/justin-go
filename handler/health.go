package handler

import (
	"fmt"
	"net/http"
)

type Health struct{}

func (h *Health) Live(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Healthy")
}

func (h *Health) Ready(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Ready")
}

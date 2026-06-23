package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shivamxverma/studentapi/internal/types"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the todo API")
}

func createTodosHandler(w http.ResponseWriter, r *http.Request) {
	var req types.CreateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(req)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /", createTodosHandler)

	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("server error:" , err)
	}
}

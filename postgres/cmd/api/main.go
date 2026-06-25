package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shivamxverma/postgres/internal/db"
)

type Handler struct {
	db *pgxpool.Pool
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Task API Running")
}

func (h *Handler) createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received task with title: %s\n", req.Title)

	var id int

	err = h.db.QueryRow(
			r.Context(), 
			"INSERT INTO tasks (title) VALUES ($1) RETURNING id",
			req.Title,
	).Scan(&id)

	if err != nil {
		fmt.Fprintln(w, "Error on creating task")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	response := map[string]any{
		"id": 		id,
		"status":  "success",
		"message": "Task created successfully",
		"title":   req.Title,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	pool, err := db.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	handler := Handler{db: pool}

	defer pool.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /create-task", handler.createTodoHandler)

	server := &http.Server{
		Addr: ":8080",
		Handler : mux,
	}

	err = server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
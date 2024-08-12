package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
	UserName    string `json:"userName"`
}

func main() {
	db := initDb()
	defer db.Close()

	server := http.NewServeMux()

	tasksController := &TasksController{DB: db}

	server.HandleFunc("GET /tasks", tasksController.Get)
	server.HandleFunc("POST /tasks", tasksController.Post)
	server.HandleFunc("PUT /tasks/{task_id}", tasksController.Put)
	server.HandleFunc("DELETE /tasks/{task_id}", tasksController.Delete)

	log.Println("Running server on port 3001")
	log.Fatal(http.ListenAndServe(":3001", server))
}

func initDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./std-lib-crud.db")
	if err != nil {
		log.Fatal("Failed to connect with sqlite3: ", err)
	}

	_, err = db.Exec("create table if not exists tasks (id integer, title text, description text, isDone int, userName text);")
	if err != nil {
		log.Fatal("Failed to create tasks table: ", err)
	}

	return db
}

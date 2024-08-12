package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Controller interface {
	Post(wr http.ResponseWriter, req *http.Request)
	Get(wr http.ResponseWriter, req *http.Request)
	Put(wr http.ResponseWriter, req *http.Request)
	Delete(wr http.ResponseWriter, req *http.Request)
}

type TasksController struct {
	DB *sql.DB
}

func (tc *TasksController) Delete(wr http.ResponseWriter, req *http.Request) {
	task_id := req.PathValue("task_id")
	if task_id == "" {
		http.Error(wr, "invalid task id", http.StatusBadRequest)
	}

	_, err := tc.DB.Exec("delete from tasks where id = ?", task_id)
	if err != nil {
		http.Error(wr, "can't delete this", http.StatusBadRequest)
	}

	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
}

func (tc *TasksController) Put(wr http.ResponseWriter, req *http.Request) {
	task_id := req.PathValue("task_id")
	if task_id == "" {
		http.Error(wr, "invalid task id", http.StatusBadRequest)
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("error reading body: ", err)
		http.Error(wr, "Could not read the body", http.StatusInternalServerError)
	}

	var todo Todo
	err = json.Unmarshal(bodyBytes, &todo)
	if err != nil {
		log.Println("Error decoding todo: ", err)
		http.Error(wr, "can't decode this shit", http.StatusBadRequest)
		return
	}

	int_task_id, err := strconv.Atoi(task_id)
	if err != nil {
		log.Println("deu ruim na hora de converter")
		http.Error(wr, "deu ruim", http.StatusInternalServerError)
	}

	todo.Id = int_task_id

	defer req.Body.Close()

	_, err = tc.DB.Exec(
		"update tasks set title = ?, description = ?, isDone = ?, userName = ? where id = ?",
		todo.Title, todo.Description, todo.IsDone, todo.UserName, todo.Id)

	if err != nil {
		log.Fatal(err)
		http.Error(wr, "Error inserting one of those todo", http.StatusBadRequest)
	}

	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated)
	json.NewEncoder(wr).Encode(todo)
}

func (tc *TasksController) Post(wr http.ResponseWriter, req *http.Request) {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("error reading body: ", err)
		http.Error(wr, "Could not read the body", http.StatusInternalServerError)
	}
	// just puttin this here

	var todos []Todo
	err = json.Unmarshal(bodyBytes, &todos)

	if err != nil {
		log.Println("Error decoding todo: ", err)
		http.Error(wr, "can't decode this shit", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < len(todos); i++ {
		todos[i].Id = rand.Intn(1000000000000) + 1

		_, err = tc.DB.Exec("insert into tasks values (?, ?, ?, ?, ?);", todos[i].Id, todos[i].Title, todos[i].Description, todos[i].IsDone, todos[i].UserName)

		if err != nil {
			log.Fatal(err)
			http.Error(wr, "Error inserting one of those todo", http.StatusBadRequest)
		}
	}

	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated)
	json.NewEncoder(wr).Encode(todos)
}

func (tc *TasksController) Get(wr http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println("Failed to parse form")
	}

	userName := req.FormValue("username")

	rows, err := tc.DB.Query("select * from tasks", userName)
	if err != nil {
		log.Println(err)
		http.Error(wr, "bad things happens", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.IsDone, &todo.UserName)

		if err != nil {
			http.Error(wr, "oops, something went wrong bro", http.StatusInternalServerError)
			return
		}

		todos = append(todos, todo)
	}

	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(todos)
}

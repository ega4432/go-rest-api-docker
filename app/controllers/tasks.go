package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go-rest-api-docker/database"

	"github.com/gorilla/mux"
)

const tableName string = "tasks"

type Task struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hell get %s:%s", vars["id"], vars["name"])
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hell getAll")
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if task.Title == "" || task.Body == "" {
		log.Println("invalid request format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stmt, err := database.Db.Prepare(fmt.Sprintf("INSERT INTO %s (title, body, created_at, updated_at) VALUES (?, ?, ?, ?)", tableName))

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer stmt.Close()

	now := time.Now().Format("2006-01-02 15:04:05")
	task.CreatedAt = now
	task.UpdatedAt = now
	res, err := stmt.Exec(task.Title, task.Body, task.CreatedAt, task.UpdatedAt)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	task.Id = int(id)

	var buf bytes.Buffer
	je := json.NewEncoder(&buf)

	if err = je.Encode(&task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprint(w, buf.String())
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hell Update id: %s", vars["id"])
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Hell delete id : %s", vars["is"])
}

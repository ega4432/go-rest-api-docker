package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go-rest-api-docker/database"

	"github.com/gorilla/mux"
)

type Task struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
	fmt.Fprintf(w, "Hell create")

	title := r.FormValue("title")
	body := r.FormValue("body")

	stmt, err := database.Db.Prepare("")

	fmt.Println(stmt)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	data := Task{
		Title:     title,
		Body:      body,
		CreatedAt: now,
		UpdatedAt: now,
	}

	var buf bytes.Buffer
	j := json.NewEncoder(&buf)
	if err := j.Encode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, _ = fmt.Fprint(w, buf.String())
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

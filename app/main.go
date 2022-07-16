package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-rest-api-docker/controllers"

	"github.com/gorilla/mux"
)

type Data struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	data := Data{
		Status:  200,
		Message: "Welcome to the Go REST API!",
	}

	var buf bytes.Buffer
	j := json.NewEncoder(&buf)

	if err := j.Encode(&data); err != nil {
		log.Fatal(err)
	}

	res := buf.String()

	log.Println(res)
	_, err := fmt.Fprint(w, res)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler)

	// tasks route
	s := r.PathPrefix("/controllers").Subrouter()
	s.HandleFunc("", controllers.GetAllHandler).Methods("GET")
	s.HandleFunc("", controllers.CreateHandler).Methods("POST")
	s.HandleFunc("/{id}", controllers.GetHandler).Methods("GET")
	s.HandleFunc("/{id}", controllers.UpdateHandler).Methods("PUT")
	s.HandleFunc("/{id}", controllers.DeleteHandler).Methods("DELETE")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

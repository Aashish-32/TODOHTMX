package routes

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	hostname string = "localhost:9000"
	port     string = ":9000"
)

func sendtodos(w http.ResponseWriter, r *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/template.html")

	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)

}

func marktodo(w http.ResponseWriter, r *http.Request) {

}
func createtodo(w http.ResponseWriter, r *http.Request) {

}

func deletetodo(w http.ResponseWriter, r *http.Request) {

}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", marktodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deletetodo).Methods("DELETE")
	mux.HandleFunc("/create", createtodo).Methods("POST")
	mux.HandleFunc("/", index)

	fmt.Printf("listening to port: %v\n", port)
	err := http.ListenAndServe(hostname, mux)

	if err != nil {
		log.Fatal(err)

	}

}

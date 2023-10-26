package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Aashish32/htmx/model"
	"github.com/gorilla/mux"
)

const (
	hostname string = "localhost:9000"
	port     string = ":9000"
)

func sendtodos(w http.ResponseWriter) {
	todos, err := model.GetallTodos()
	if err != nil {
		fmt.Println("could not fetch all todos from database", err)
		return
	}

	tmpl, err := template.ParseFiles("./templates/template.html")

	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.ExecuteTemplate(w, "Todos", todos); err != nil {
		log.Fatal(err)
	}

}

func index(w http.ResponseWriter, r *http.Request) {

	todos, err := model.GetallTodos()
	if err != nil {
		fmt.Println("could not fetch all todos from database", err)
		return
	}

	tmpl, err := template.ParseFiles("./templates/template.html")

	if err != nil {
		log.Fatal(err)
	}
	if err := tmpl.Execute(w, todos); err != nil {
		log.Fatal(err)
	}

}

func marktodo(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		fmt.Println("cant convert id to uint64")
	}
	err = model.MarkCompleted(id)
	if err != nil {
		fmt.Println("could not mark the task as completed")
	}
	sendtodos(w)
}
func createtodo(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Println("Error parsing form")
	}
	err := model.CreateTodo(r.FormValue("todo"))
	if err != nil {
		fmt.Println("cant create todo")
	}
	sendtodos(w)

}

func deletetodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		fmt.Println("cant convert id to uint64")
	}
	err = model.Deletetodo(id)
	if err != nil {
		fmt.Println("could not delete the task")
	}
	sendtodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", marktodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deletetodo).Methods("DELETE")
	mux.HandleFunc("/create", createtodo).Methods("POST")

	fmt.Printf("listening to port: %v\n", port)
	err := http.ListenAndServe(hostname, mux)

	if err != nil {
		log.Fatal(err)

	}

}

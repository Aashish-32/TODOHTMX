package model

import "fmt"

type Todo struct {
	Id        uint64 `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
}

func CreateTodo(todo string) error {
	statement := "insert into todos(todo , completed) values ($1, $2); "

	if _, err := db.Query(statement, todo, false); err != nil {
		fmt.Println("couldnot query db to get all todos", err)
		return err

	}

	return nil
}
func GetallTodos() ([]Todo, error) {
	todos := []Todo{}
	statement := "Select * from todos"
	rows, err := db.Query(statement)

	if err != nil {
		fmt.Println("couldnot query db to get all todos", err)
		return todos, err

	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var completed bool
		var id uint64

		err := rows.Scan(&id, &title, &completed)
		if err != nil {
			return todos, err
		}
		todo := Todo{
			Id:        id,
			Todo:      title,
			Completed: completed,
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodo(Id uint64) (Todo, error) {
	todo := Todo{}
	statement := "Select * from todos where Id=$1"

	row, err := db.Query(statement, Id)
	if err != nil {
		return todo, err
	}

	for row.Next() {
		var title string
		var completed bool
		var Id uint64

		err := row.Scan(&Id, &title, &completed)
		if err != nil {
			return todo, err
		}

		todo = Todo{
			Id:        Id,
			Todo:      title,
			Completed: completed,
		}
	}
	return todo, err
}

func MarkCompleted(Id uint64) error {

	todo, err := GetTodo(Id)
	if err != nil {
		return err
	}
	statement := "update todos set completed=$2 where Id=$1;"

	_, err = db.Exec(statement, Id, !todo.Completed)
	if err != nil {
		return err

	}
	return nil
}

func Deletetodo(Id uint64) error {
	statement := "delete from todos where Id=$1"

	_, err := db.Exec(statement, Id)
	return err
}

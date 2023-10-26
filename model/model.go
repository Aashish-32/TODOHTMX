package model

import "fmt"

type Todo struct {
	id        uint64 `json:"id"`
	todo      string `json:"todo"`
	completed bool   `json:"completed"`
}

func CreateTodo(todo string) error {
	statement := "insert into todos(todo , completed) values ($1, $2); "

	_, err := db.Query(statement, todo, false)

	return err
}
func GetallTodos() ([]Todo, error) {
	todos := []Todo{}
	statement := "Select * from todos"
	rows, err := db.Query(statement)

	if err != nil {
		fmt.Println(err)
		return todos, err

	}
	defer rows.Close()
	for rows.Next() {
		var title string
		var completed bool
		var id uint64

		if err := rows.Scan(&id, &title, &completed); err != nil {
			return todos, err
		}
		todo := Todo{
			id:        id,
			todo:      title,
			completed: completed,
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodo(id uint64) (Todo, error) {
	todo := Todo{}
	statement := "Select * from todos where id==$1"

	row, err := db.Query(statement, id)
	if err != nil {
		return todo, err
	}

	for row.Next() {
		var title string
		var completed bool
		var id uint64

		err := row.Scan(&id, &title, &completed)
		if err != nil {
			return todo, err
		}

		todo = Todo{
			id:        id,
			todo:      title,
			completed: completed,
		}
	}
	return todo, err
}

func MarkCompleted(id uint64) error {

	if todo, err := GetTodo(id); err != nil {
		return err
	} else {

		statement := "update todos set completed =$2 where id = $1 "

		_, err := db.Exec(statement, id, !todo.completed)
		if err != nil {
			return err

		}
	}
	return nil
}

func delete(id uint64) error {
	statement := "delete from todos where id==$1"

	_, err := db.Exec(statement, id)
	return err
}

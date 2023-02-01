package models

import (
	"github.com/diogo-braz/go-api-postgresql/database"
)

func InsertTodo(todo Todo) (id int64, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}

func GetTodo(id int64) (todo Todo, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT id, title, description, done FROM todos WHERE id = $1`, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}

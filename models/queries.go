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

func GetAllTodos() (todos []Todo, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT id, title, description, done FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}

func UpdateTodo(id int64, todo Todo) (int64, error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title = $2, description = $3, done = $4 WHERE id = $1`,
		id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func DeleteTodo(id int64) (int64, error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

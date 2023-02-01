package database

import (
	"database/sql"
	"fmt"

	"github.com/diogo-braz/go-api-postgresql/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled",
		conf.Hostname, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}

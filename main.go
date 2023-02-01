package main

import (
	"fmt"
	"net/http"

	"github.com/diogo-braz/go-api-postgresql/configs"
	"github.com/diogo-braz/go-api-postgresql/handlers"
	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.CreateTodo)
	r.Put("/{id}", handlers.UpdateTodo)
	r.Delete("/{id}", handlers.DeleteTodo)
	r.Get("/", handlers.ListTodos)
	r.Get("/{id}", handlers.GetTodo)

	http.ListenAndServe(fmt.Sprintf(":%v", configs.GetServerPort()), r)
}

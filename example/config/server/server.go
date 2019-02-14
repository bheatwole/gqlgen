package main

import (
	"log"
	"net/http"

	todo "github.com/bheatwole/gqlgen/example/config"
	"github.com/bheatwole/gqlgen/handler"
)

func main() {
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

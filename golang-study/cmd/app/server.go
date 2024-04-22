package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Takenari-Yamamoto/golang-study/generated/graph"
	"github.com/Takenari-Yamamoto/golang-study/internal/repository"
	"github.com/Takenari-Yamamoto/golang-study/internal/resolver"
	"github.com/Takenari-Yamamoto/golang-study/internal/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
		TaskService: taskService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

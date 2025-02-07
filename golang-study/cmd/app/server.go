package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Takenari-Yamamoto/golang-study/generated/graph"
	"github.com/Takenari-Yamamoto/golang-study/internal/repository"
	"github.com/Takenari-Yamamoto/golang-study/internal/resolver"
	"github.com/Takenari-Yamamoto/golang-study/internal/service"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// TODO: config.go とかに切り出したいかも
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to init database: ", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect db: ", err)
		return
	}

	log.Default().Println("success to connect db!!")

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{
		TaskService: taskService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

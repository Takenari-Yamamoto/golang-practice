package main

import (
	"context"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/graph"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/graph/loader"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/repository"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	todoRepo := repository.NewTodoRepository()
	userLoader := loader.NewUserLoader()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		TodoRepo:   todoRepo,
		UserLoader: userLoader,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", DataLoaderMiddleware(srv)) // Middlewareを追加

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type contextKey string

func DataLoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := loader.NewUserLoader()

		ctx := context.WithValue(r.Context(), contextKey("userLoader"), loader)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

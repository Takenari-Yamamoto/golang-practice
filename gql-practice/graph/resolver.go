package graph

import (
	"github/Takenari-Yamamoto/golang-practice/gql-practice/repository"

	"github.com/graph-gophers/dataloader"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoRepo   *repository.TodoRepository
	UserLoader *dataloader.Loader
}

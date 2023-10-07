package graph

import "github/Takenari-Yamamoto/golang-practice/gql-practice/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRepo *repository.UserRepository
	todoRepo *repository.TodoRepository
}

package repository

import "github/Takenari-Yamamoto/golang-practice/gql-practice/domain"

func ListAllTodos() []*domain.Todo {
	res := []domain.Todo{
		{
			ID:     "1",
			Text:   "todo1",
			Done:   false,
			UserId: "1",
		},
		{
			ID:     "2",
			Text:   "todo2",
			Done:   false,
			UserId: "2",
		},
		{
			ID:     "3",
			Text:   "todo3",
			Done:   false,
			UserId: "3",
		},
	}
	var todos []*domain.Todo
	for _, v := range res {
		todos = append(todos, &v)
	}
	return todos
}

package repository

import (
	"fmt"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/domain"
)

func ListAllTodos() []*domain.Todo {
	var res []*domain.Todo
	for i := 1; i <= 100; i++ {
		res = append(res, &domain.Todo{
			ID:     fmt.Sprintf("todo%d", i),
			Text:   "todo" + fmt.Sprintf("%d", i),
			Done:   false,
			UserId: domain.GenerateUserId(),
		})
	}

	return res
}

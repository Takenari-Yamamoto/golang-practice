package repository

import (
	"fmt"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/domain"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/graph/model"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (repo *TodoRepository) ListAllTodos() ([]*model.Todo, error) {
	var res []*domain.Todo
	for i := 1; i <= 100; i++ {
		res = append(res, &domain.Todo{
			ID:     fmt.Sprintf("todo%d", i),
			Text:   "todo" + fmt.Sprintf("%d", i),
			Done:   false,
			UserId: domain.GenerateUserId(),
		})
	}
	var todos []*model.Todo
	for _, v := range res {
		todos = append(todos, &model.Todo{
			ID:     v.ID,
			Text:   v.Text,
			Done:   v.Done,
			UserID: v.UserId,
		})
	}
	return todos, nil
}

package repository

import (
	"context"
	"fmt"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/domain"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/graph/model"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/pkg/errors"

	"github.com/samber/lo"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (repo *TodoRepository) ListAllTodos() ([]*model.Todo, error) {
	var todos []*domain.Todo
	for i := 1; i <= 100; i++ {
		todos = append(todos, &domain.Todo{
			ID:     fmt.Sprintf("todo%d", i),
			Text:   "todo" + fmt.Sprintf("%d", i),
			Done:   false,
			UserId: domain.GenerateUserId(),
		})
	}

	return lo.Map(todos, func(t *domain.Todo, _ int) *model.Todo {
		return &model.Todo{
			ID:     t.ID,
			Text:   t.Text,
			Done:   t.Done,
			UserID: t.UserId,
		}
	}), nil
}

func (repo *TodoRepository) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	var errorDetails []errors.ErrorDetail

	if input.Text == "" {
		errorDetails = append(errorDetails, errors.ErrorDetail{
			Message:   "テキストを入力してください",
			Attribute: "text",
		})
	}

	if input.UserID == "" {
		errorDetails = append(errorDetails, errors.ErrorDetail{
			Message:   "ユーザーIDを入力してください",
			Attribute: "userId",
		})
	}

	if len(errorDetails) > 0 {
		return nil, errors.NewCustomError(errors.GqlError{
			Message:     "invalid input",
			Code:        errors.BAD_REQUEST,
			UserMessage: "入力値が不正です",
			Details:     &errorDetails,
		})
	}

	return &model.Todo{
		ID:     "todo1",
		Text:   "todo1",
		Done:   false,
		UserID: domain.GenerateUserId(),
	}, nil
}

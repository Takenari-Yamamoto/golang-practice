package repository

import "github.com/Takenari-Yamamoto/golang-study/internal/domain"

type ITaskRepository interface {
	ListAll() ([]*domain.Task, error)
}

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) ListAll() ([]*domain.Task, error) {
	tasks := []*domain.Task{
		{ID: "1", Title: "Title1", Description: "Description1"},
		{ID: "2", Title: "Title2", Description: "Description2"},
		{ID: "3", Title: "Title3", Description: "Description3"},
	}
	return tasks, nil
}

package service

import (
	"context"

	"github.com/Takenari-Yamamoto/golang-study/internal/domain"
	"github.com/Takenari-Yamamoto/golang-study/internal/repository"
)

type TaskService struct {
	taskRepository *repository.TaskRepository
}

func NewTaskService(
	taskRepository *repository.TaskRepository,
) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (s *TaskService) ListAll(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := s.taskRepository.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	task, err := s.taskRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) Create(ctx context.Context, task *domain.Task) error {
	if err := s.taskRepository.Create(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s *TaskService) Update(ctx context.Context, task *domain.Task) error {
	if err := s.taskRepository.Update(ctx, task); err != nil {
		return err
	}

	return nil
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
	if err := s.taskRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

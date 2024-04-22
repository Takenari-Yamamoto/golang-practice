package service

import (
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

func (s *TaskService) ListAll() ([]*domain.Task, error) {
	tasks, err := s.taskRepository.ListAll()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

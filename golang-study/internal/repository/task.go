package repository

import (
	"context"
	"database/sql"

	"github.com/Takenari-Yamamoto/golang-study/generated/database/models"
	"github.com/Takenari-Yamamoto/golang-study/internal/domain"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

type ITaskRepository interface {
	ListAll(ctx context.Context) ([]*domain.Task, error)
	FindByID(ctx context.Context, id string) (*domain.Task, error)
	Create(ctx context.Context, task *domain.Task) error
	Update(ctx context.Context, task *domain.Task) error
	Delete(ctx context.Context, id string) error
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(
	db *sql.DB,
) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) ListAll(ctx context.Context) ([]*domain.Task, error) {
	tasks, err := models.Tasks().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var result []*domain.Task
	for _, task := range tasks {
		result = append(result, &domain.Task{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description.String,
		})
	}
	return result, nil

}

func (r *TaskRepository) FindByID(ctx context.Context, id string) (*domain.Task, error) {
	task, err := models.Tasks(models.TaskWhere.ID.EQ(id)).One(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description.String,
	}, nil
}

func (r *TaskRepository) Create(ctx context.Context, task *domain.Task) error {
	taskToInsert := &models.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: null.NewString(task.Description, task.Description != ""),
	}
	if err := taskToInsert.Insert(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Update(ctx context.Context, task *domain.Task) error {
	taskToUpdate := &models.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: null.NewString(task.Description, task.Description != ""),
	}
	if _, err := taskToUpdate.Update(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	taskToDelete := &models.Task{
		ID: id,
	}
	if _, err := taskToDelete.Delete(ctx, r.db); err != nil {
		return err
	}
	return nil
}

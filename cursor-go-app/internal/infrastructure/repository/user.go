package repository

import (
	"database/sql"

	"github.com/Takenari-Yamamoto/golang-practice/cursor-go-app/internal/domain/model"
	"github.com/Takenari-Yamamoto/golang-practice/cursor-go-app/internal/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id int) (*model.User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)
	var user model.User
	err := row.Scan(&user.ID, &user.Name)
	return &user, err
}

package repository

import (
	"github.com/Takenari-Yamamoto/golang-practice/cursor-go-app/internal/domain/model"
)

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
}

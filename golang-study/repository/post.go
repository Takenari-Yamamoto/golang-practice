package repository

import "github/Takenari-Yamamoto/golang-practice/golang-study/domain"

type IPostRepository interface {
	FindAll() ([]domain.Post, error)
}

type PostRepository struct{}

func NewPostRepository() IPostRepository {
	return PostRepository{}
}

func (r PostRepository) FindAll() ([]domain.Post, error) {
	posts := []domain.Post{
		{ID: 1, Title: "Hello, World!"},
		{ID: 2, Title: "Hello, Golang!"},
	}
	return posts, nil
}

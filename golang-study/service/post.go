package service

import (
	"github/Takenari-Yamamoto/golang-practice/golang-study/domain"
	"github/Takenari-Yamamoto/golang-practice/golang-study/repository"
)

type IPostService interface {
	FindAll() ([]*domain.Post, error)
}

type PostService struct {
	PostRepository repository.IPostRepository
}

func NewPostService(r repository.IPostRepository) IPostService {
	return &PostService{PostRepository: r}
}

func (s *PostService) FindAll() ([]*domain.Post, error) {
	posts, err := s.PostRepository.FindAll()
	if err != nil {
		return nil, err
	}
	res := make([]*domain.Post, 0, len(posts))
	for i := range posts {
		res = append(res, &posts[i])
	}
	return res, nil
}

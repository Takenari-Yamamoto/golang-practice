package repository

import (
	"fmt"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/domain"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/graph/model"

	"github.com/samber/lo"
)

type UserRepository struct{}

func (repo *UserRepository) ListAllUsers() []*domain.User {
	res := []*domain.User{
		{
			ID:   domain.USER_ID_TKNR1216,
			Name: "たけなり",
		},
		{
			ID:   domain.USER_ID_TAKASHI003,
			Name: "たかし",
		},
		{
			ID:   domain.USER_ID_TARO115,
			Name: "太郎",
		},
		{
			ID:   domain.USER_ID_HANAKO43,
			Name: "花子",
		},
		{
			ID:   domain.USER_ID_YAMADA234,
			Name: "山田",
		},
	}
	return res
}

func (repo *UserRepository) GetUserByID(id string) (*model.User, error) {

	all := repo.ListAllUsers()
	res, ok := lo.Find(all, func(user *domain.User) bool {
		return user.ID == id
	})

	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return &model.User{
		ID:   res.ID,
		Name: res.Name,
	}, nil

}

func (repo *UserRepository) GetUsersByIDs(ids []string) ([]*model.User, error) {

	allUsers := repo.ListAllUsers()
	users := lo.Filter(allUsers, func(user *domain.User, _ int) bool {
		return lo.Contains(ids, user.ID)
	})

	return lo.Map(users, func(user *domain.User, _ int) *model.User {
		return &model.User{
			ID:   user.ID,
			Name: user.Name,
		}
	}), nil
}

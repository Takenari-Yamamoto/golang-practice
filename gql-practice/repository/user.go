package repository

import (
	"fmt"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/domain"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/graph/model"
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
	fmt.Println("ユーザーを取得します", id)
	all := repo.ListAllUsers()
	var res domain.User
	for _, v := range all {
		if v.ID == id {
			res = *v
		}
	}

	return &model.User{
		ID:   res.ID,
		Name: res.Name,
	}, nil

}

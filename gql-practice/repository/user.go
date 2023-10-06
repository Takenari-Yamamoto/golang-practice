package repository

import "github/Takenari-Yamamoto/golang-practice/gql-practice/domain"

func ListAllUsers() []*domain.User {
	res := []domain.User{
		{
			ID:   "1",
			Name: "user1",
		},
		{
			ID:   "2",
			Name: "user2",
		},
		{
			ID:   "3",
			Name: "user3",
		},
	}
	var users []*domain.User
	for _, v := range res {
		users = append(users, &v)
	}
	return users
}

func GetUserByID(id string) *domain.User {
	all := ListAllUsers()
	var res domain.User
	for _, v := range all {
		if v.ID == id {
			res = *v
		}
	}
	return &res
}

package repository

import (
	"context"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

type UserRepository struct {
	spanner *spanner.Client
}

func NewUserRepository(
	spanner *spanner.Client,
) *UserRepository {
	return &UserRepository{
		spanner: spanner,
	}
}

// TODO: これは切り出す
type User struct {
	ID        string `spanner:"ID"`
	Name      string `spanner:"Name"`
	CreatedBy string `spanner:"CreatedBy"`
}

func (r *UserRepository) ListAll(ctx context.Context) ([]User, error) {
	stmt := spanner.Statement{
		SQL: `SELECT Name FROM Users`,
	}
	iter := r.spanner.Single().Query(ctx, stmt)
	defer iter.Stop()

	// TODO: この辺の処理は共通化したい
	var users []User
	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var user User
		if err := row.ToStruct(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}

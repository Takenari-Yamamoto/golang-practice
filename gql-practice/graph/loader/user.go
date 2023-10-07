package loader

import (
	"context"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/graph/model"
	"github/Takenari-Yamamoto/golang-practice/gql-practice/repository"

	"github.com/graph-gophers/dataloader"
)

func NewUserLoader() *dataloader.Loader {
	return dataloader.NewBatchedLoader(batchFunc, dataloader.WithClearCacheOnBatch())
}

func batchFunc(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	userRepo := &repository.UserRepository{}

	var ids []string
	for _, key := range keys {
		ids = append(ids, key.String())
	}

	users, err := userRepo.GetUsersByIDs(ids)

	results := make([]*dataloader.Result, len(keys))
	if err != nil {
		for i := range keys {
			results[i] = &dataloader.Result{Error: err}
		}
		return results
	}

	userMap := make(map[string]*model.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	for i, key := range keys {
		id := key.String()
		results[i] = &dataloader.Result{Data: userMap[id]}
	}

	return results
}

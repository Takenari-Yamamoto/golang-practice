package main

import (
	"context"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	client, err := spanner.NewClient(ctx, "projects/sample-pro-395702/instances/spanner-sample-instance/databases/spanner-sample-db",
		option.WithCredentialsFile("credentials/key.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// データの挿入
	err = insertData(ctx, client)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	// データのクエリ
	err = queryData(ctx, client)
	if err != nil {
		log.Fatalf("Failed to query data: %v", err)
	}
}

func insertData(ctx context.Context, client *spanner.Client) error {
	mutations := []*spanner.Mutation{
		spanner.InsertOrUpdate("Users",
			[]string{"UserId", "UserName", "Email"},
			[]interface{}{"12345", "John Doe", "john.doe@example.com"}),
	}

	_, err := client.Apply(ctx, mutations)
	return err
}

func queryData(ctx context.Context, client *spanner.Client) error {
	stmt := spanner.Statement{SQL: `SELECT UserId, UserName, Email FROM Users`}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		var userId, userName, email string
		if err := row.Columns(&userId, &userName, &email); err != nil {
			return err
		}

		log.Printf("UserId: %s, UserName: %s, Email: %s", userId, userName, email)
	}
	return nil
}

package emulator

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
)

var MaxWait = 10 * time.Second

const (
	projectID      = "test-project-id"
	instanceID     = "test-instance"
	displayName    = "Test Instance"
	instanceConfig = "projects/" + projectID + "/instanceConfigs/emulator-config"
)

func CreateDatabase(ctx context.Context, schema []string) string {

	if os.Getenv("SPANNER_EMULATOR_HOST") == "" {
		log.Fatalf("Spanner emulator is not running")
	}

	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create database admin client: %v", err)
	}
	defer adminClient.Close()

	database := "test-database"
	instance := "projects/" + projectID + "/instances/" + instanceID

	op, err := adminClient.CreateDatabase(ctx, &databasepb.CreateDatabaseRequest{
		Parent:          instance,
		CreateStatement: "CREATE DATABASE `" + database + "`",
		ExtraStatements: []string{
			`CREATE TABLE Users (
				ID STRING(MAX) NOT NULL,
				Name STRING(MAX) NOT NULL,
				CreatedBy STRING(MAX) NOT NULL,
			) PRIMARY KEY (ID)`,
		},
		// 本当は下記のようにしたいが、フォーマットがうまく行かず、エラーになるので修正が必要
		// ExtraStatements: schema,
	})
	if err != nil {
		log.Fatalf("adminClient.CreateDatabase: %v", err)
	}

	if _, err := op.Wait(ctx); err != nil {
		log.Fatalf("op.Wait: %v", err)
	}
	return instance + "/databases/" + database
}

func SetupDatabase(schema []string, data ...*spanner.Mutation) *spanner.Client {

	ctx := context.Background()
	database := CreateDatabase(ctx, schema)
	client, err := spanner.NewClient(ctx, database)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}

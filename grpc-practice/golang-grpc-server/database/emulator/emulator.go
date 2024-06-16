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

	// grpcOptions := []option.ClientOption{
	// 	option.WithEndpoint("localhost:9010"),
	// 	option.WithoutAuthentication(),
	// 	option.WithGRPCDialOption(grpc.WithBlock()),
	// 	option.WithGRPCDialOption(grpc.WithConnectParams(grpc.ConnectParams{
	// 		Backoff:           backoff.DefaultConfig,
	// 		MinConnectTimeout: 5 * time.Second,
	// 	})),
	// }

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
		ExtraStatements: schema,
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

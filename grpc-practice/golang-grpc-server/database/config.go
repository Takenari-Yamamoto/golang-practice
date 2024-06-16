package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/spanner"
	admin "cloud.google.com/go/spanner/admin/database/apiv1"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

const (
	Project  = "test-project"
	Instance = "test-instance"
	Name     = "test-database"
	Id       = "test-database"
)

// NewSpannerClient creates a new Spanner client.
func NewSpannerClient(ctx context.Context) (*spanner.Client, error) {

	spannerEndpoint := "localhost:9010"
	database := fmt.Sprintf("projects/%s/instances/%s/databases/%s", Project, Instance, Name)

	client, err := spanner.NewClient(ctx, database, option.WithEndpoint(spannerEndpoint), option.WithoutAuthentication())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client, nil
}

// NewDatabaseAdminClient creates a new DatabaseAdmin client.
func NewDatabaseAdminClient(ctx context.Context) (*admin.DatabaseAdminClient, error) {

	grpcOptions := []option.ClientOption{
		option.WithEndpoint("localhost:9010"),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithConnectParams(grpc.ConnectParams{
			Backoff:           backoff.DefaultConfig,
			MinConnectTimeout: 5 * time.Second,
		})),
	}
	adminClient, err := admin.NewDatabaseAdminClient(ctx, grpcOptions...)
	if err != nil {
		log.Fatalf("Failed to create database admin client: %v", err)
		return nil, err
	}
	defer adminClient.Close()

	return adminClient, nil
}

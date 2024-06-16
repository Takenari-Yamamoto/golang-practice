package util

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/option"
)

const (
	project  = "test-project"
	instance = "test-instance"
	database = "test-database"
)

func NewSpannerClient(ctx context.Context) (*spanner.Client, error) {

	spannerEndpoint := "localhost:9010"
	database := fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, database)

	client, err := spanner.NewClient(ctx, database, option.WithEndpoint(spannerEndpoint), option.WithoutAuthentication())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client, nil
}

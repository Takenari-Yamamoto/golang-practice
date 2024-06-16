package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Takenari-Yamamoto/golang-grpc-server/database"
)

func main() {
	ctx := context.Background()
	dbClient, err := database.NewSpannerClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer dbClient.Close()

	fmt.Println("Spanner client is created successfully")

}

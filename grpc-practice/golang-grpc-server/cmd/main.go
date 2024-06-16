package main

import (
	"context"
	"fmt"
	"log"

	util "github.com/Takenari-Yamamoto/golang-grpc-server/util"
)

func main() {
	ctx := context.Background()
	dbClient, err := util.NewSpannerClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer dbClient.Close()

	fmt.Println("Spanner client is created successfully")

}

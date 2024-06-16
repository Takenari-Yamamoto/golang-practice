package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/spanner"
	"github.com/Takenari-Yamamoto/golang-grpc-server/repository"
)

// TODO: 共通化
const (
	projectID   = "test-project-id"
	instanceID  = "test-instance"
	displayName = "Test Instance"
	databaseID  = "test-database"
	port        = "8080"
)

func main() {
	fmt.Println("main app server is running...")

	ctx := context.Background()

	db := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectID, instanceID, databaseID)
	client, err := spanner.NewClient(ctx, db)
	if err != nil {
		log.Fatalf("Failed to create client %v", err)
		return
	}
	defer client.Close()
	log.Default().Printf("Spanner client is created successfully: %v", client.ClientID())

	repository := repository.NewUserRepository(client)

	// http handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, HTTP server!")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		res, err := repository.ListAll(ctx)
		if err != nil {
			log.Fatalf("failed to get users: %v", err)
			return
		}

		for _, v := range res {
			fmt.Fprintf(w, "ID: %s, Name: %s, CreatedBy: %s\n", v.ID, v.Name, v.CreatedBy)
		}

		b, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("failed to marshal json: %v", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	})

	// start server
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("failed to start server: %v", err)
		return
	}
	log.Default().Printf("server is running on port %s", port)

}

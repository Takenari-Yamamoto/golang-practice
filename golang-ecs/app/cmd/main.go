package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received!")
		w.Write([]byte("Hello, Takenari samansa chuchu!!!!!!!!"))
	})
	fmt.Println("Server started")

	http.ListenAndServe(":8080", nil)
}

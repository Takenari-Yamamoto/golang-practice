package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received!")
		w.Write([]byte("Hello, Take-chan!!!!!!!!"))
	})
	fmt.Println("Server started")

	http.ListenAndServe(":8080", nil)
}

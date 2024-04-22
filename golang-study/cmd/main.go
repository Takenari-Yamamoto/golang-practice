package main

import (
	"github/Takenari-Yamamoto/golang-practice/golang-study/internal/repository"
	"github/Takenari-Yamamoto/golang-practice/golang-study/internal/service"
)

func main() {

	postRepo := repository.NewPostRepository()
	postService := service.NewPostService(postRepo)

	posts, err := postService.FindAll()
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		println(post.ID, post.Title)
	}

}

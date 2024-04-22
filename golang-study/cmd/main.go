package main

import (
	"github/Takenari-Yamamoto/golang-practice/golang-study/repository"
	"github/Takenari-Yamamoto/golang-practice/golang-study/service"
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

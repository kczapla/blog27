package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func handleRequests() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.Exec("PRAGMA foreign_keys = ON")
	postRepository := NewRepository(db)
	postService := NewService(postRepository)

	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)

	commentRepository := NewCommentRepository(db)
	commentService := NewCommentService(commentRepository)

	tagRepository := NewTagRepository(db)
	tagService := NewTagService(tagRepository)

	router := mux.NewRouter().StrictSlash(true)
	RegisterHandlers(postService, router)
	RegisterUserHandlers(userService, router)
	RegisterCommentHandlers(commentService, router)
	RegisterTagHandlers(tagService, router)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func initialMigration() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Tag{})
}

func main() {
	initialMigration()
	handleRequests()
}

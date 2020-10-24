package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func handleRequests() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }
    postRepository := NewRepository(db)
    postService := NewService(postRepository)

    userRepository := NewUserRepository(db)
    userService := NewUserService(userRepository)

    router := mux.NewRouter().StrictSlash(true)
    RegisterHandlers(postService, router)
    RegisterUserHandlers(userService, router)

    log.Fatal(http.ListenAndServe(":8081", router))
}


func initialMigration() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    db.AutoMigrate(&User{}, &Post{})
}


func main() {
    initialMigration()
    handleRequests()
}

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
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/posts", AllPosts).Methods("GET")
    myRouter.HandleFunc("/post", NewPost).Methods("POST")
    myRouter.HandleFunc("/post/{id}", DeletePost).Methods("DELETE")
    //myRouter.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
    //myRouter.HandleFunc("/user", newUser).Methods("POST")
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}



func initialMigration() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    db.AutoMigrate(&Post{})
}


func main() {
    initialMigration()
    handleRequests()
}

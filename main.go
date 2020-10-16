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
    //myRouter.HandleFunc("/users/{name}/{email}", updateUser).Methods("PUT")
    //myRouter.HandleFunc("/user", newUser).Methods("POST")
    myRouter.HandleFunc("/users", AllUsers).Methods("GET")
    myRouter.HandleFunc("/user/{id}", GetUser).Methods("GET")
    myRouter.HandleFunc("/user", NewUser).Methods("POST")
    myRouter.HandleFunc("/user/{id}", UpdateUser).Methods("POST")
    myRouter.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")
    myRouter.HandleFunc("/user/{id}/posts", AllUserPosts).Methods("GET")
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}



func initialMigration() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    db.AutoMigrate(&Post{})
    db.AutoMigrate(&User{})
    db.AutoMigrate(&UserPost{})
}


func main() {
    initialMigration()
    handleRequests()
}

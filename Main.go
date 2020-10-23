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
    repository := NewRepository(db)
    service := NewService(repository)
    router := mux.NewRouter().StrictSlash(true)

    RegisterHandlers(service, router)
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

package main

import (
    "fmt"
    "log"
    "net/http"

    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func handleRequests() {
    myRouter := NewRouter()
    log.Fatal(http.ListenAndServe(":8081", myRouter))
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

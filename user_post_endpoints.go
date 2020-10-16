package main

import (
    "fmt"
    //"log"
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func AllPosts(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    var posts []Post
    db.Find(&posts)
    json.NewEncoder(w).Encode(posts)
}


func NewPost(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    reqBody, _ := ioutil.ReadAll(r.Body)
    var post Post
    json.Unmarshal(reqBody, &post)
    db.Create(&post)
}


func DeletePost(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    key := vars["id"]

    db.Delete(&Post{}, key)

}
//
//
//func updateUser(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "update user endpoint hit")
//}

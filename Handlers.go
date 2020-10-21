package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    var users []User
    db.Find(&users)
    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }
    vars := mux.Vars(r)
    key := vars["id"]

    var user User
    db.First(&user, key)
    json.NewEncoder(w).Encode(user)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    reqBody, _ := ioutil.ReadAll(r.Body)
    var user User
    json.Unmarshal(reqBody, &user)
    db.Create(&user)
    json.NewEncoder(w).Encode(user)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Printf("id %v", key)
    db.Delete(&User{}, key)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }
    vars := mux.Vars(r)
    key := vars["id"]

    reqBody, _ := ioutil.ReadAll(r.Body)
    var userInRequest User
    json.Unmarshal(reqBody, &userInRequest)

    var userInDB User
    db.First(&userInDB, key)

    if userInRequest.Name != "" {
        userInDB.Name = userInRequest.Name
    }

    if userInRequest.Email != "" {
        userInDB.Email = userInRequest.Email
    }
    db.Save(&userInDB)
    json.NewEncoder(w).Encode(&userInDB)
}


func CreatePost(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    reqBody, _ := ioutil.ReadAll(r.Body)
    var post Post

    json.Unmarshal(reqBody, &post)

    db.Create(&post)
    json.NewEncoder(w).Encode(post)
}


func GetPost(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    key := vars["id"]

    var post Post
    db.First(&post, key)

    json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    var post []Post
    db.Find(&post)

    json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }

    vars := mux.Vars(r)
    key := vars["id"]
    var currentPost Post
    db.First(&currentPost, key)

    reqBody, _ := ioutil.ReadAll(r.Body)
    var incomingPost Post
    json.Unmarshal(reqBody, &incomingPost)

    if (incomingPost.Title != "") {
        currentPost.Title = incomingPost.Title
    }

    if (incomingPost.Content != "") {
        currentPost.Content = incomingPost.Content
    }

    db.Save(currentPost)

    json.NewEncoder(w).Encode(currentPost)
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

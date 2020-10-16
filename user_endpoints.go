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

    db.Delete(&User{}, key)

}

func AllUserPosts(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("failed to connect database")
    }
    vars := mux.Vars(r)
    key := vars["id"]

    var userPosts []UserPost
    db.First(&userPosts, key)
    json.NewEncoder(w).Encode(userPosts)
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

    if userInRequest.Nickname != "" {
        userInDB.Nickname = userInRequest.Nickname
    }

    if userInRequest.Email != "" {
        userInDB.Email = userInRequest.Email
    }
    db.Save(&userInDB)
    json.NewEncoder(w).Encode(&userInDB)
}

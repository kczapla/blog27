package main


import (
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

func RegisterUserHandlers(service UserService, router *mux.Router) {
    userResource := userResource{service}
    router.HandleFunc("/users/{id}", userResource.get).Methods("GET")
    router.HandleFunc("/users", userResource.create).Methods("POST")
    router.HandleFunc("/users/{id}", userResource.delete).Methods("DELETE")
    router.HandleFunc("/users/{id}", userResource.update).Methods("PATCH")
    router.HandleFunc("/users", userResource.query).Methods("GET")
}

type userResource struct {
    service UserService
}

func (res userResource) get(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    post, _ := res.service.Get(key)
    json.NewEncoder(w).Encode(post)
}

func (res userResource) create(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var userCreateRequest UserCreateRequest
    json.Unmarshal(reqBody, &userCreateRequest)

    err := res.service.Create(userCreateRequest)
    if err != nil {
        http.Error(w, "Post not created", http.StatusForbidden)
    }
}

func (res userResource) delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    err := res.service.Delete(key)

    if err != nil {
        http.Error(w, "Post not deleted", http.StatusForbidden)
    }
}

func (res userResource) update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    reqBody, _ := ioutil.ReadAll(r.Body)
    var userUpdateRequest UserUpdateRequest

    json.Unmarshal(reqBody, &userUpdateRequest)
    err := res.service.Update(key, userUpdateRequest)
    if err != nil {
        http.Error(w, "Post not updated", http.StatusForbidden)
        return
    }
}

func (res userResource) query(w http.ResponseWriter, r *http.Request) {
    posts, err := res.service.Query()
    if err != nil {
        http.Error(w, "Can't querry", http.StatusForbidden)
    }

    json.NewEncoder(w).Encode(posts)
}

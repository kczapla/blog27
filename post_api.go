package main


import (
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

func RegisterHandlers(service Service, router *mux.Router) {
    resource := resource{service}
    router.HandleFunc("/posts/{id}", resource.get).Methods("GET")
    router.HandleFunc("/posts", resource.create).Methods("POST")
    router.HandleFunc("/posts/{id}", resource.delete).Methods("DELETE")
    router.HandleFunc("/posts/{id}", resource.update).Methods("PATCH")
    router.HandleFunc("/posts", resource.query).Methods("GET")
}

type resource struct {
    service Service
}

func (res resource) get(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    post, _ := res.service.Get(key)
    json.NewEncoder(w).Encode(post)
}

func (res resource) create(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var postCreateRequest PostCreateRequest
    json.Unmarshal(reqBody, &postCreateRequest)

    err := res.service.Create(postCreateRequest)
    if err != nil {
        http.Error(w, "Post not created", http.StatusForbidden)
    }
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    err := res.service.Delete(key)

    if err != nil {
        http.Error(w, "Post not deleted", http.StatusForbidden)
    }
}

func (res resource) update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    reqBody, _ := ioutil.ReadAll(r.Body)
    var postUpdateRequest PostUpdateRequest

    json.Unmarshal(reqBody, &postUpdateRequest)
    err := res.service.Update(key, postUpdateRequest)
    if err != nil {
        http.Error(w, "Post not updated", http.StatusForbidden)
        return
    }
}

func (res resource) query(w http.ResponseWriter, r *http.Request) {
    posts, err := res.service.Query()
    if err != nil {
        http.Error(w, "Can't querry", http.StatusForbidden)
    }

    json.NewEncoder(w).Encode(posts)
}

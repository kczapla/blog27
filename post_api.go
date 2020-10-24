package main


import (
    "io/ioutil"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

func RegisterHandlers(service Service, router *mux.Router) {
    postResource := postResource{service}
    router.HandleFunc("/posts/{id}", postResource.get).Methods("GET")
    router.HandleFunc("/posts", postResource.create).Methods("POST")
    router.HandleFunc("/posts/{id}", postResource.delete).Methods("DELETE")
    router.HandleFunc("/posts/{id}", postResource.update).Methods("PATCH")
    router.HandleFunc("/posts", postResource.query).Methods("GET")
}

type postResource struct {
    service Service
}

func (res postResource) get(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    post, _ := res.service.Get(key)
    json.NewEncoder(w).Encode(post)
}

func (res postResource) create(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var postCreateRequest PostCreateRequest
    json.Unmarshal(reqBody, &postCreateRequest)

    err := res.service.Create(postCreateRequest)
    if err != nil {
        http.Error(w, "Post not created", http.StatusForbidden)
    }
}

func (res postResource) delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]
    err := res.service.Delete(key)

    if err != nil {
        http.Error(w, "Post not deleted", http.StatusForbidden)
    }
}

func (res postResource) update(w http.ResponseWriter, r *http.Request) {
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

func (res postResource) query(w http.ResponseWriter, r *http.Request) {
    posts, err := res.service.Query()
    if err != nil {
        http.Error(w, "Can't querry", http.StatusForbidden)
    }

    json.NewEncoder(w).Encode(posts)
}

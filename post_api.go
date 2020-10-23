package main


import (
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

func RegisterHandlers(service Service, router *mux.Router) {
    resource := resource{service}
    router.HandleFunc("/posts/{id}", resource.get).Methods("GET")
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

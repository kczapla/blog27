package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func RegisterTagHandlers(service TagService, router *mux.Router) {
	userResource := tagResource{service}
	router.HandleFunc("/tags/{id}", userResource.get).Methods("GET")
	router.HandleFunc("/tags", userResource.create).Methods("POST")
	router.HandleFunc("/tags/{id}", userResource.delete).Methods("DELETE")
	router.HandleFunc("/tags/{id}", userResource.update).Methods("PATCH")
	router.HandleFunc("/tags", userResource.query).Methods("GET")
}

type tagResource struct {
	service TagService
}

type tagQuery struct {
	Name string
}

func (res tagResource) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	post, _ := res.service.Get(key)
	json.NewEncoder(w).Encode(post)
}

func (res tagResource) create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userCreateRequest TagCreateRequest
	json.Unmarshal(reqBody, &userCreateRequest)

	err := res.service.Create(userCreateRequest)
	if err != nil {
		http.Error(w, "Post not created", http.StatusForbidden)
	}
}

func (res tagResource) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	err := res.service.Delete(key)

	if err != nil {
		http.Error(w, "Post not deleted", http.StatusForbidden)
	}
}

func (res tagResource) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userUpdateRequest TagUpdateRequest

	json.Unmarshal(reqBody, &userUpdateRequest)
	err := res.service.Update(key, userUpdateRequest)
	if err != nil {
		http.Error(w, "Post not updated", http.StatusForbidden)
		return
	}
}

func (res tagResource) query(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	var userQuery userQuery
	decodeErr := decoder.Decode(&userQuery, r.URL.Query())

	if decodeErr != nil {
		http.Error(w, "Can't convert querry", http.StatusForbidden)
		return
	}

	var users Tags
	var user Tag
	var queryErr error

	if userQuery.Name != "" {
		user, queryErr = res.service.QueryTagByName(userQuery.Name)
		if queryErr != nil {
			http.Error(w, "Can't querry user", http.StatusForbidden)
		}
		json.NewEncoder(w).Encode(user)
	} else {
		users, queryErr = res.service.QueryAllTags()
		if queryErr != nil {
			http.Error(w, "Can't querry", http.StatusForbidden)
		}
		json.NewEncoder(w).Encode(users)
	}
}

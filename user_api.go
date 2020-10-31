package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
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

type userQuery struct {
	Name string
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
	decoder := schema.NewDecoder()
	var userQuery userQuery
	decodeErr := decoder.Decode(&userQuery, r.URL.Query())

	if decodeErr != nil {
		http.Error(w, "Can't convert querry", http.StatusForbidden)
		return
	}

	var users Users
	var user User
	var queryErr error

	if userQuery.Name != "" {
		user, queryErr = res.service.QueryUserByName(userQuery.Name)
		if queryErr != nil {
			http.Error(w, "Can't querry user", http.StatusForbidden)
		}
		json.NewEncoder(w).Encode(user)
	} else {
		users, queryErr = res.service.QueryAllUsers()
		if queryErr != nil {
			http.Error(w, "Can't querry", http.StatusForbidden)
		}
		json.NewEncoder(w).Encode(users)
	}
}

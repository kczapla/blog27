package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	if len(r.URL.Query()) == 0 {
		posts, err := res.queryAllPosts()
		if err != nil {
			http.Error(w, "Can't querry", http.StatusForbidden)
		}
		json.NewEncoder(w).Encode(posts)
		return
	}

	posts, err := res.queryAllUserPosts(r.URL.Query().Get("userId"))
	if err != nil {
		http.Error(w, "Can't querry", http.StatusForbidden)
	}
	json.NewEncoder(w).Encode(posts)
	// var posts Posts
	// var err error

	// if len(r.URL.Query()) == 0 {
	// 	fmt.Printf("first")
	// 	posts, err = res.queryAllPosts()
	// } else if r.URL.Query().Get("userId") == "" {
	// 	fmt.Printf("second")
	// 	posts, err = res.queryAllUserPosts(r.URL.Query().Get("userId"))
	// }

	// if err != nil {
	// 	http.Error(w, "Can't querry", http.StatusForbidden)
	// 	return
	// }

	// json.NewEncoder(w).Encode(posts)
}

func (res postResource) queryAllPosts() (Posts, error) {
	posts, err := res.service.QueryAllPosts()
	return posts, err
}

func (res postResource) queryAllUserPosts(userId string) (Posts, error) {
	uid, conversionErr := strconv.ParseUint(userId, 10, 64)
	if conversionErr != nil {
		fmt.Println("error converting user id to int")
		return Posts{}, conversionErr
	}

	return res.service.QueryAllUserPosts(uint(uid))
}

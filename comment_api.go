package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func RegisterCommentHandlers(service CommentService, router *mux.Router) {
	resource := commentResource{service}
	router.HandleFunc("/comments/{id}", resource.get).Methods("GET")
	router.HandleFunc("/comments", resource.create).Methods("POST")
	router.HandleFunc("/comments/{id}", resource.delete).Methods("DELETE")
	router.HandleFunc("/comments/{id}", resource.update).Methods("PATCH")
	router.HandleFunc("/comments", resource.query).Methods("GET")
}

type commentResource struct {
	service CommentService
}

type commentQuery struct {
	UserID uint
	PostID uint
}

func (res commentResource) get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	post, _ := res.service.Get(key)
	json.NewEncoder(w).Encode(post)
}

func (res commentResource) create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var commentCreateRequest CommentCreateRequest
	json.Unmarshal(reqBody, &commentCreateRequest)

	err := res.service.Create(commentCreateRequest)
	if err != nil {
		http.Error(w, "Comment not created", http.StatusForbidden)
	}
}

func (res commentResource) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	err := res.service.Delete(key)

	if err != nil {
		http.Error(w, "Comment not deleted", http.StatusForbidden)
	}
}

func (res commentResource) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var commentUpdateRequest CommentUpdateRequest

	json.Unmarshal(reqBody, &commentUpdateRequest)
	err := res.service.Update(key, commentUpdateRequest)
	if err != nil {
		http.Error(w, "Comment not updated", http.StatusForbidden)
		return
	}
}

func (res commentResource) query(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	var commentQuery commentQuery
	decodeErr := decoder.Decode(&commentQuery, r.URL.Query())

	if decodeErr != nil {
		http.Error(w, "Can't convert querry", http.StatusForbidden)
		return
	}

	var comments Comments
	var queryErr error

	if commentQuery.UserID != 0 {
		comments, queryErr = res.service.QueryUserComments(commentQuery.UserID)
	} else if commentQuery.PostID != 0 {
		comments, queryErr = res.service.QueryPostComments(commentQuery.PostID)
	} else {
		comments, queryErr = res.service.QueryAll()
	}

	if queryErr != nil {
		http.Error(w, "Can't querry", http.StatusForbidden)
	}

	json.NewEncoder(w).Encode(comments)

}

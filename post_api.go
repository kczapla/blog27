package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func RegisterHandlers(service Service, router *mux.Router) {
	s := router.PathPrefix("/posts").Subrouter()
	postResource := postResource{service}
	s.HandleFunc("", postResource.query).Methods("GET")
	s.HandleFunc("", postResource.create).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", postResource.get).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", postResource.delete).Methods("DELETE")
	s.HandleFunc("/{id:[0-9]+}", postResource.update).Methods("PATCH")
	s.HandleFunc("/{id:[0-9]+}/tags", postResource.queryTags).Methods("GET")
	s.HandleFunc("/{postId:[0-9]+}/tags/{tagId:[0-9]+}", postResource.addTag).Methods("POST")
	s.HandleFunc("/{postId:[0-9]+}/tags/{tagId:[0-9]+}", postResource.deleteTag).Methods("DELETE")
	s.HandleFunc("/tags", postResource.queryTaggedPosts).Methods("GET")
}

type postResource struct {
	service Service
}

type postQuery struct {
	UserID uint
}

type taggedPostQuerry struct {
	Name []string
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
	decoder := schema.NewDecoder()
	var postQuery postQuery
	decodeErr := decoder.Decode(&postQuery, r.URL.Query())

	if decodeErr != nil {
		http.Error(w, "Can't convert querry", http.StatusForbidden)
		return
	}

	var posts Posts
	var queryErr error

	if postQuery.UserID != 0 {
		posts, queryErr = res.service.QueryAllUserPosts(postQuery.UserID)
	} else {
		posts, queryErr = res.service.QueryAllPosts()
	}

	if queryErr != nil {
		http.Error(w, "Can't querry", http.StatusForbidden)
	}

	json.NewEncoder(w).Encode(posts)
}

func (res postResource) queryTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		http.Error(w, "Can't convert passed id", http.StatusForbidden)
		return
	}
	post, _ := res.service.QueryPostTags(uint(id))
	json.NewEncoder(w).Encode(post)
}

func (res postResource) addTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, postIdConversionErr := strconv.ParseUint(vars["postId"], 10, 64)
	tagId, tagIdConversionErr := strconv.ParseUint(vars["tagId"], 10, 64)

	if postIdConversionErr != nil || tagIdConversionErr != nil {
		http.Error(w, "Can not process passed ids", http.StatusForbidden)
		return
	}

	addTagErr := res.service.AddTag(uint(postId), uint(tagId))

	if addTagErr != nil {
		http.Error(w, "Tag was not added", http.StatusForbidden)
	}
}

func (res postResource) deleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, postIdConversionErr := strconv.ParseUint(vars["postId"], 10, 64)
	tagId, tagIdConversionErr := strconv.ParseUint(vars["tagId"], 10, 64)

	if postIdConversionErr != nil || tagIdConversionErr != nil {
		http.Error(w, "Can not process passed ids", http.StatusForbidden)
		return
	}

	deleteTagErr := res.service.DeleteTag(uint(postId), uint(tagId))

	if deleteTagErr != nil {
		http.Error(w, "Tag was not deleted", http.StatusForbidden)
	}
}

func (res postResource) queryTaggedPosts(w http.ResponseWriter, r *http.Request) {
	tagsIdsStrings := r.URL.Query()["tagId"]
	var tagsIds []uint

	for _, stringTagId := range tagsIdsStrings {
		tagId, err := strconv.ParseUint(stringTagId, 10, 64)
		if err != nil {
			continue
		}
		tagsIds = append(tagsIds, uint(tagId))
	}

	posts, queryError := res.service.QueryPostsWithTags(tagsIds)
	if queryError != nil {
		http.Error(w, "Can't find posts", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

package main

import (
	"strconv"
)

type Service interface {
	Get(id string) (Post, error)
	Create(postCreateRequest PostCreateRequest) error
	Delete(id string) error
	Update(id string, postUpdateRequest PostUpdateRequest) error
	QueryAllPosts() (Posts, error)
	QueryAllUserPosts(userId uint) (Posts, error)
	QueryPostTags(postId uint) (Tags, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) service {
	return service{repository}
}

type PostCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
}

type PostUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (s service) Get(id string) (Post, error) {
	post, err := s.repository.Get(id)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (s service) Create(postCreateRequest PostCreateRequest) error {
	var post Post
	post.Title = postCreateRequest.Title
	post.Content = postCreateRequest.Content
	post.UserID = postCreateRequest.UserID

	err := s.repository.Create(post)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Update(id string, postUpdateRequest PostUpdateRequest) error {
	post, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	if postUpdateRequest.Title != "" {
		post.Title = postUpdateRequest.Title
	}

	if postUpdateRequest.Content != "" {
		post.Content = postUpdateRequest.Content
	}

	err = s.repository.Update(post)
	if err != nil {
		return err
	}
	return nil
}

func (s service) QueryAllPosts() (Posts, error) {
	posts, err := s.repository.Query(Post{})
	if err != nil {
		return Posts{}, err
	}
	return posts, nil
}

func (s service) QueryAllUserPosts(userId uint) (Posts, error) {
	posts, err := s.repository.Query(Post{UserID: userId})
	if err != nil {
		return Posts{}, nil
	}
	return posts, nil
}

func (s service) QueryPostTags(postId uint) (Tags, error) {
	post, postQueryError := s.repository.Get(strconv.FormatUint(uint64(postId), 10))
	if postQueryError != nil {
		return Tags{}, postQueryError
	}

	tags, err := s.repository.QueryTags(post)
	if err != nil {
		return Tags{}, err
	}

	if len(tags) == 0 {
		return Tags{}, nil
	}

	return tags, nil
}

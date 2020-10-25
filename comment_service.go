package main

type CommentService interface {
	Get(id string) (Comment, error)
	Create(commentCreateRequest CommentCreateRequest) error
	Delete(id string) error
	Update(id string, commentUpdateRequest CommentUpdateRequest) error
	Query() (Comments, error)
}

type commentService struct {
	repository CommentRepository
}

func NewCommentService(repository CommentRepository) commentService {
	return commentService{repository}
}

type CommentCreateRequest struct {
	UserID  uint   `json:"userId"`
	PostID  uint   `json:"postId"`
	Content string `json:"content"`
}

type CommentUpdateRequest struct {
	Content string `json:"content"`
}

func (s commentService) Get(id string) (Comment, error) {
	comment, err := s.repository.Get(id)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}

func (s commentService) Create(commentCreateRequest CommentCreateRequest) error {
	var comment Comment
	comment.UserID = commentCreateRequest.UserID
	comment.PostID = commentCreateRequest.PostID
	comment.Content = commentCreateRequest.Content

	err := s.repository.Create(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s commentService) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s commentService) Update(id string, commentUpdateRequest CommentUpdateRequest) error {
	comment, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	if commentUpdateRequest.Content != "" {
		comment.Content = commentUpdateRequest.Content
	}

	err = s.repository.Update(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s commentService) Query() (Comments, error) {
	comment, err := s.repository.Query()
	if err != nil {
		return Comments{}, err
	}
	return comment, nil
}

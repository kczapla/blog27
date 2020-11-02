package main

type TagService interface {
	Get(id string) (Tag, error)
	Create(tagCreateRequest TagCreateRequest) error
	Delete(id string) error
	Update(id string, tagUpdateRequest TagUpdateRequest) error
	QueryAllTags() (Tags, error)
	QueryTagByName(tagName string) (Tag, error)
}

type tagService struct {
	repository TagRepository
}

func NewTagService(repository TagRepository) tagService {
	return tagService{repository}
}

type TagCreateRequest struct {
	Name string `json:"name"`
}

type TagUpdateRequest struct {
	Name string `json:"name"`
}

func (s tagService) Get(id string) (Tag, error) {
	tag, err := s.repository.Get(id)
	if err != nil {
		return Tag{}, err
	}

	return tag, nil
}

func (s tagService) Create(tagCreateRequest TagCreateRequest) error {
	var tag Tag
	tag.Name = tagCreateRequest.Name

	err := s.repository.Create(tag)
	if err != nil {
		return err
	}
	return nil
}

func (s tagService) Delete(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s tagService) Update(id string, tagUpdateRequest TagUpdateRequest) error {
	tag, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	if tagUpdateRequest.Name != "" {
		tag.Name = tagUpdateRequest.Name
	}

	err = s.repository.Update(tag)
	if err != nil {
		return err
	}
	return nil
}

func (s tagService) QueryAllTags() (Tags, error) {
	tags, err := s.repository.Query(Tag{})
	if err != nil {
		return Tags{}, err
	}
	return tags, nil
}

func (s tagService) QueryTagByName(tagName string) (Tag, error) {
	tags, err := s.repository.Query(Tag{Name: tagName})
	if err != nil {
		return Tag{}, err
	}

	if len(tags) == 0 {
		return Tag{}, nil
	}

	return tags[0], nil
}

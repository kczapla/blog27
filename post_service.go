package main


type Service interface {
    Get(id string) (Post, error)
}


type service struct {
    repository Repository
}

func NewService(repository Repository) service {
    return service{repository}
}


func (s service) Get(id string) (Post, error) {
    post, err := s.repository.Get(id)
    if err != nil {
        return Post{}, err
    }

    return post, nil
}

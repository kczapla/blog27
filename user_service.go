package main

type UserService interface {
    Get(id string) (User, error)
    Create(userCreateRequest UserCreateRequest) error
    Delete(id string) error
    Update(id string, userUpdateRequest UserUpdateRequest) error
    Query() (Users, error)
}


type userService struct {
    repository UserRepository
}

func NewUserService(repository UserRepository) userService {
    return userService{repository}
}

type UserCreateRequest struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

type UserUpdateRequest struct {
    Name string `json:"name"`
    Email string `json:"email"`
}


func (s userService) Get(id string) (User, error) {
    user, err := s.repository.Get(id)
    if err != nil {
        return User{}, err
    }

    return user, nil
}

func (s userService) Create(userCreateRequest UserCreateRequest) error {
    var user User
    user.Name = userCreateRequest.Name
    user.Email = userCreateRequest.Email

    err := s.repository.Create(user)
    if err != nil {
        return err
    }
    return nil
}

func (s userService) Delete(id string) error {
    err := s.repository.Delete(id)
    if err != nil {
        return err
    }
    return nil
}

func (s userService) Update(id string, userUpdateRequest UserUpdateRequest) error {
    user, err := s.repository.Get(id)
    if err != nil {
        return err
    }

    if userUpdateRequest.Name != "" {
        user.Name = userUpdateRequest.Name
    }

    if userUpdateRequest.Email != "" {
        user.Email = userUpdateRequest.Email
    }

    err = s.repository.Update(user)
    if err != nil {
        return err
    }
    return nil
}

func (s userService) Query() (Users, error) {
    user, err := s.repository.Query()
    if err != nil {
        return Users{}, err
    }
    return user, nil
}


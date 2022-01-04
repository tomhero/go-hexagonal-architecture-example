package service

import "gofiber/repository"

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) Register(userRequest UserRequest) (*UserResponse, error) {
	return nil, nil
}

func (s userService) Login(username string, password string) (*UserResponse, error) {
	return nil, nil
}

package service

import (
	"gofiber/errs"
	"gofiber/repository"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s userService) Register(userRequest UserRequest) (*UserResponse, error) {
	fixRole := "user"

	// NOTE : Encrypt Password
	hashedPassword, err := hashPassword(userRequest.Password)
	if err != nil {
		return nil, err
	}

	user := repository.User{
		Username:    userRequest.Username,
		Password:    hashedPassword,
		Role:        fixRole,
		Customer_id: userRequest.CustomerID,
	}

	err = s.userRepo.Create(&user)
	if err != nil {
		return nil, errs.NewUnexpectedError()
	}

	response := UserResponse{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
		Token:    "fake jwt token",
	}
	return &response, nil
}

func (s userService) Login(username string, password string) (*UserResponse, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, errs.NewValidationError("incorrect username or password")
	}

	if !checkPasswordHash(password, user.Password) {
		return nil, errs.NewValidationError("incorrect username or password")
	}

	response := UserResponse{
		Role:  user.Role,
		Token: "fake jwt token",
	}
	return &response, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

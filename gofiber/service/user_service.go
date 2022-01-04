package service

import (
	"gofiber/errs"
	"gofiber/logs"
	"gofiber/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
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
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	token, err := generateToken(user.Customer_id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := UserResponse{
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
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

	token, err := generateToken(user.Customer_id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := UserResponse{
		Role:  user.Role,
		Token: token,
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

func generateToken(customerID int) (string, error) {
	cliams := jwt.StandardClaims{
		Issuer:    strconv.Itoa(customerID),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)

	token, err := jwtToken.SignedString([]byte(viper.GetString("app.jwtSecret")))
	if err != nil {
		return "", err
	}

	return token, nil
}

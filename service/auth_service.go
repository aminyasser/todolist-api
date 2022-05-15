package service

import (
	"errors"
	"log"

	"github.com/aminyasser/todo-list/entity/model"
	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/repository"
	"github.com/mashingan/smapping"
)

type AuthService interface {
	RegisterUser(user model.UserRequest) (*response.UserResponse, error)
}
type authService struct {
	repository repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *authService {
	return &authService{repo}
}

func (auth *authService) RegisterUser(user model.UserRequest) (*response.UserResponse, error) {
	emailExists := auth.repository.Exists(user.Email)
	if emailExists {
		return nil, errors.New("email already exists")
	}
	userMapped := model.User{}
	err := smapping.FillStruct(&userMapped, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil , err
	}

	 userRegisterd := auth.repository.Insert(userMapped)
	 res := response.NewUserResponse(userRegisterd)
	 return &res , nil
}
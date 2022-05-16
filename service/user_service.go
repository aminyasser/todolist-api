package service

import (
	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/repository"
)


type UserService interface {
	FindUser(id string) (*response.ProfileResponse , error)
}
type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo}
}

func (user *userService) FindUser(id string) (*response.ProfileResponse , error) {
	profile , err := user.repository.FindBy("id" , id)
	if err != nil {
		return nil , err
	}
    res := response.NewProfileResponse(profile)
	return  &res , nil
}

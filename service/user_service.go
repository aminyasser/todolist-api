package service

import "github.com/aminyasser/todo-list/repository"


type UserService interface {

}
type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo}
}

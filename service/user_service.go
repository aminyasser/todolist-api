package service

import (


	"github.com/aminyasser/todo-list/entity/model"
	"github.com/aminyasser/todo-list/entity/request"
	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/repository"
	"github.com/mashingan/smapping"
)


type UserService interface {
	FindUser(id string) (*response.ProfileResponse , error)
	UpdateUser( profile request.ProfileUpdate) (*response.ProfileResponse , error) 
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

func (user *userService) UpdateUser( profile request.ProfileUpdate) (*response.ProfileResponse , error)  {
	userModel := model.User{}
	err := smapping.FillStruct(&userModel, smapping.MapFields(&profile))
    userModel.ID = profile.ID
	if err != nil {
		return nil, err
	}

	updatedUser := user.repository.Update(userModel)
	
	res := response.NewProfileResponse(updatedUser)
	return &res, nil
}

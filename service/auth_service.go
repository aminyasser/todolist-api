package service

import (
	"errors"
	"log"

	"github.com/aminyasser/todo-list/entity/model"
	"github.com/aminyasser/todo-list/entity/request"
	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(user request.UserRegister) (*response.UserResponse, error)
	VerifyUser( string,  string)  (*response.UserResponse , error) 
}
type authService struct {
	repository repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *authService {
	return &authService{repo}
}

func (auth *authService) RegisterUser(user request.UserRegister) (*response.UserResponse, error) {
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

// VerifyUser takes the credintials and return the response user
func (c *authService) VerifyUser(email string, password string) (*response.UserResponse , error) {
	user, err := c.repository.FindBy("email" , email)
	if err != nil {
		println(err.Error())
		return nil , err
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return nil , errors.New("failed, check your credential")
	}
    res := response.NewUserResponse(user)
	return &res , nil
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
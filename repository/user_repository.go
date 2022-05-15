package repository

import (
	"fmt"

	"github.com/aminyasser/todo-list/entity/model"
	"gorm.io/gorm"
)

type UserRepository interface {
Insert( model.User) model.User
Exists( string) bool
}

type userRepository struct {
	connection  *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (auth *userRepository) Insert(user model.User) model.User {

    res:=  auth.connection.Create(&user)
    fmt.Println(res.RowsAffected , user)
	return user
}

func (auth *userRepository) Exists(email string) bool {
    var user model.User
	res := auth.connection.Where("email = ?", email).Take(&user)

	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
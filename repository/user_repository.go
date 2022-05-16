package repository

import (
	"log"

	"github.com/aminyasser/todo-list/entity/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
FindBy(string , string) (model.User, error)
Insert( model.User) model.User
Update( model.User) model.User
Exists( string) bool
}

type userRepository struct {
	connection  *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (auth *userRepository) FindBy(colm string , value string) (model.User, error) {
	var user model.User
	res := auth.connection.Where(colm +" = ?", value).Take(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}


func (auth *userRepository) Insert(user model.User) model.User {
	user.Password = hashPassword([]byte(user.Password))
    auth.connection.Create(&user)
	return user
}

func (auth *userRepository) Update(user model.User) model.User {
	auth.connection.Where("id = ?", user.ID).Updates(&user)
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

func hashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

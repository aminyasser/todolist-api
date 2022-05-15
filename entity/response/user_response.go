package response

import "github.com/aminyasser/todo-list/entity/model"

type UserResponse struct {
	ID    uint  `json:"id"`
	Name string  `json:"name"`
	Email string  `json:"email"`
	Token string  `json:"token"`
}


func NewUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

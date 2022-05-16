package request



type UserRegister struct {
	Name string  `json:"name" bindding:"required"`
	Email string  `json:"email" bindding:"required"`
	Password string  `json:"password" bindding:"required"`
}

type UserLogin struct {
	Email string  `json:"email" bindding:"required"`
	Password string  `json:"password" bindding:"required"`
}

type ProfileUpdate struct {
	ID    uint      `json:"id" bindding:"required"`
	Name string  `json:"name" bindding:"required"`
	Email string  `json:"email" bindding:"required"`
}
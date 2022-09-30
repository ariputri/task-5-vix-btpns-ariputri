package models

// LoginModel is a model that's used by the user when POST from /login url
type LoginModel struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}

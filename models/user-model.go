package models

import "time"

// UserUpdate is used by user when PUT update profile
type UserUpdateModel struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
	//Photo     string `json:"photo" form:"photo" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

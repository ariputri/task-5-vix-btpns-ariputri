package app

import "time"

// User struct represents users table in database
type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id" binding:"required"`
	Username string `gorm:"type:varchar(255)" json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(256);UNIQUE" json:"email" binding:"required,email"`
	Password string `gorm:"->;<-; not null" json:"-" binding:"required"`
	Token    string `gorm:"-" json:"token, omitempty"`
	//Photos    *[]Photo  `gorm:"foreignkey:PhotoID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"photos,omitempty"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

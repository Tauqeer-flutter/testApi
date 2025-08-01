package response

import (
	"time"
)

type SuccessAuthResponse struct {
	Status  bool     `json:"status" default:"true"`
	Message string   `json:"message"`
	Token   string   `json:"token"`
	User    UserData `json:"user"`
}

type UserData struct {
	Id         uint      `gorm:"primaryKey;autoIncrement" json:"id" required:"false"`
	FirstName  string    `gorm:"column:firstName;size:20;not null;" json:"firstName" validate:"required"`
	LastName   string    `gorm:"column:lastName;size:20;not null" json:"lastName" validate:"required"`
	Email      string    `gorm:"size:100;not null;unique" json:"email" validate:"required,email"`
	Age        int       `gorm:"size:3;not null" json:"age" validate:"required"`
	IsVerified bool      `gorm:"default:true" json:"is_verified" required:"false"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at" required:"false"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at" required:"false"`
}

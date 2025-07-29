package dtos

import "time"

type User struct {
	Id         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname  string    `gorm:"size:20;not null" json:"firstname"`
	Lastname   string    `gorm:"size:20;not null" json:"lastname"`
	Email      string    `gorm:"size:20;not null;unique" json:"email"`
	Password   string    `gorm:"size:100;not null" json:"password"`
	Age        int       `gorm:"size:3;not null" json:"age"`
	IsVerified bool      `gorm:"default:true" json:"is_verified"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

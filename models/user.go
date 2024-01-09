package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Username  string         `json:"username" gorm:"unique"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Email     string         `json:"email"`
	IsAdmin   bool           `json:"isAdmin" gorm:"default:false"`
	Password  string         `json:"-" gorm:"->;<-;not null" gorm:"->:false;<-:create"`
	IsActive  bool           `json:"isActive" gorm:"default:true"`
}

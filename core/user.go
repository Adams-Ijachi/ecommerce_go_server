package core

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Email       *string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Age         uint8      `gorm:"type:int;not null" json:"age"`
	Username    string     `gorm:"type:varchar(100);uniqueIndex" json:"username"`
	DOB         *time.Time `gorm:"type:date" json:"dob"`
	Password    string     `gorm:"not null"`
	PhoneNumber string     `gorm:"type:varchar(100);uniqueIndex" json:"phone_number"`
}

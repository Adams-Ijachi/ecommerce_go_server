package core

import (
	"encoding/json"
	"time"

	"ecom/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string  `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName    string  `gorm:"type:varchar(100);not null" json:"last_name"`
	Email       *string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Age         uint8   `gorm:"type:int;not null" json:"age"`
	Username    string  `gorm:"type:varchar(100);uniqueIndex" json:"username"`
	Password    string  `gorm:"not null" json:"password,omitempty"`
	PhoneNumber string  `gorm:"type:varchar(100);uniqueIndex" json:"phone_number"`

	UserProfile UserProfile `gorm:"foreignKey:UserID" json:"user_profile"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, _ = helpers.HashPassword(u.Password)
	return
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User // prevent recursion
	x := user(u)
	x.Password = "" // remove password from json

	return json.Marshal(x)

}

type UserProfile struct {
	gorm.Model
	UserID       uint      `gorm:"not null" json:"user_id"`
	Gender       string    `gorm:"type:varchar(100);null" `
	ProfileImage string    `gorm:"type:varchar(100);null" json:"profile_image"`
	DOB          time.Time `gorm:"type:date" json:"dob"`
	Country      string    `gorm:"type:varchar(100);not null" json:"country"`
}

type NewUser struct {
	FirstName   string  ` json:"first_name" binding:"required,alpha"`
	LastName    string  ` json:"last_name" binding:"required,alpha"`
	Email       *string ` json:"email" binding:"required,email"`
	Age         uint8   ` json:"age" binding:"required,numeric"`
	Username    string  ` json:"username" binding:"required,alphanum"`
	Password    string  ` json:"password,omitempty" binding:"required,min=8,alphanum"`
	PhoneNumber string  ` json:"phone_number" binding:"required,numeric"`

	// user profile

	Gender     string ` json:"gender" binding:"alpha"`
	ProfileImg string ` json:"profile_image" binding:"omitempty,file"`
	DOB        string `json:"dob" binding:"required,datetime=2006-01-01" `
	Country    string ` json:"country" binding:"alpha"`
}

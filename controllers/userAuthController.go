package controllers

import (
	"ecom/core"
	"ecom/datastore/postgres"
	custom_err "ecom/errors"

	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserAuthController struct {
	client *postgres.Client
}

func NewAuthController(client *postgres.Client) *UserAuthController {
	return &UserAuthController{
		client: client,
	}
}

func (u *UserAuthController) RegisterUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var userData *core.NewUser

		if err := c.ShouldBindJSON(&userData); err != nil {
			errorMessages := custom_err.GetErrors(err)

			c.JSON(400, gin.H{
				"error": errorMessages,
			})

			return
		}

		var user core.User
		u.client.DB.Where("email = ?", userData.Email).Or("username = ?", userData.Username).Or("phone_number = ?", userData.PhoneNumber).First(&user)

		if user.ID != 0 {
			c.JSON(400, gin.H{
				"error": "User already exists",
			})

			return
		}

		dob, err := time.Parse("2006-01-02", userData.DOB)
		if err != nil {
			return
		}

		//create user
		newUser := core.User{
			FirstName:   userData.FirstName,
			LastName:    userData.LastName,
			Email:       userData.Email,
			PhoneNumber: userData.PhoneNumber,
			Password:    userData.Password,
			Username:    userData.Username,
			Age:         userData.Age,
			UserProfile: core.UserProfile{
				Gender:  userData.Gender,
				Country: userData.Country,
				DOB:     dob,
			},
		}

		u.client.DB.Transaction(func(tx *gorm.DB) error {

			if err := tx.Create(&newUser).Error; err != nil {
				return err
			}

			return nil

		})

		c.JSON(200, gin.H{
			"message": "User Registered Successfully",
			"data":    newUser,
		})

		return
	}

}

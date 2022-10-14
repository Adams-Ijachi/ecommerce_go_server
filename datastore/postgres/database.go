package postgres

import (
	"ecom/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func NewClient(connection config.Configuration) *Client {

	db, err := gorm.Open(postgres.Open(connection.DatabaseURL.Url), &gorm.Config{})

	if err != nil {
		return nil
	}

	log.Println("Connected to a database")

	return &Client{
		DB: db,
	}
}

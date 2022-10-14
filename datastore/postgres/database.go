package postgres

import (
	"ecom/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewClient(connection config.Configuration) (*Client, error) {

	db, err := gorm.Open(postgres.Open(connection.DatabaseURL.Url), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("Connected to a database")

	return &Client{
		db: db,
	}, nil
}

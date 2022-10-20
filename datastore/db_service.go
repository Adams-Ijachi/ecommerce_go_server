package datastore

import "ecom/datastore/postgres"

type DbService struct {
	client *postgres.Client
}

func NewDbService(client *postgres.Client) *DbService {

	dbInstance := &DbService{
		client: client,
	}

	return dbInstance
}

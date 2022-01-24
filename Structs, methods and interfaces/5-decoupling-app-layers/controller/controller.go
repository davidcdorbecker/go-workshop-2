package controller

import (
	"decoupling/db"
	externalapi "decoupling/externalAPI"

	"fmt"
)

// Find a way to decouple the controller from the DB and the ExternalAPI layers

type Controller struct {
	db  db.DB
	api externalapi.ExternalAPI
}

func NewController(db db.DB, api externalapi.ExternalAPI) Controller {
	return Controller{
		db:  db,
		api: api,
	}
}

func (c Controller) GetUser() {
	fmt.Println("In GetUser controller!")

	c.db.GetInfoFromDB()

	c.api.GetInfoFromAPI()

	fmt.Println("Returning from GetUser controller!")
}

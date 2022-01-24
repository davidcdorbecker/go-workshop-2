package main

import (
	"decoupling/controller"
	"decoupling/db"
	externalapi "decoupling/externalAPI"
)

func main() {
	c := controller.NewController(db.DB{}, externalapi.ExternalAPI{})
	c.GetUser()
}

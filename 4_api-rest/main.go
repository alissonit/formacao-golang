package main

import (
	"fmt"

	"github.com/alissonit/api-rest/database"
	"github.com/alissonit/api-rest/routes"
)

func main() {

	// models.Personalities = []models.Personality{
	// 	{Id: 1, Name: "Alisson", History: "I'm a software developer"},
	// 	{Id: 2, Name: "John", History: "I'm a software developer too"},
	// }

	database.ConnectDatabase()

	fmt.Println("starting the server...")
	routes.HandleRequest()
}

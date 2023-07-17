package main

import (
	"github.com/alissonit/go-api-rest-gin/database"
	"github.com/alissonit/go-api-rest-gin/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}

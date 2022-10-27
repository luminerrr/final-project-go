package main

import (
	"final-project-go/database"
	"final-project-go/routers"
)

func main() {
	database.StartDB()
	router := routers.StartApp()

	router.Run(":8080")
	
}
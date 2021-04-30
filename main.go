package main

import (
	"gohiringchannels/config"
	"gohiringchannels/models"
	"gohiringchannels/routes"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	db := config.SetupDB()
	db.AutoMigrate(&models.Engineer{})
	db.AutoMigrate(&models.Enterprice{})

	r := routes.SetupRoutes(db)
	r.Run(port) //setting Port

}

package main

import (
	"OnionPractice/app/router"
	"OnionPractice/db"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	// database
	db.ConnectDatabase()

	// router
	r := router.SetupRouter()

	// start server
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		fmt.Println(err)
	}
}

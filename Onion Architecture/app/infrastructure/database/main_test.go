package database

import (
	"OnionPractice/db"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../../../.env"); err != nil {
		fmt.Println("No .env file found")
	}

	db.ConnectDatabase()

	os.Exit(m.Run())
}

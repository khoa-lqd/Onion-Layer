package integration__test

import (
	"OnionPractice/app/router"
	"OnionPractice/db"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

const testServerPort = ":9997"

func TestMain(m *testing.M) {
	// load env
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("No .env file found")
	}

	// connect database
	db.ConnectDatabase()
	// db.ConnectDataBaseCube()

	// router
	r := router.SetupRouter()

	go func() {
		if err := r.Run(testServerPort); err != nil {
			log.Fatal(err)
		}
	}()
	time.Sleep(3 * time.Second)
	os.Exit(m.Run())
}

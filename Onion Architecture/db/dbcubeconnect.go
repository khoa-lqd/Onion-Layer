package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBCube *gorm.DB
)

func ConnectDataBaseCube() {
	dbHost := os.Getenv("DB_HOST")
	dbCube := os.Getenv("DB_CUBE")
	dbUserName := os.Getenv("DB_SECURITY_USERNAME")
	dbPassword := os.Getenv("DB_SECURITY_PASSWORD")

	dsn := dbUserName + ":" + dbPassword +
		"@(" + dbHost + ")/" + dbCube + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("nrmysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database cube (nrmysql)...")
		fmt.Println(err)

		return
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		fmt.Println("Failed to connect to database cube (gorm)...")
		fmt.Println(err)

		return
	}

	// 555
	// Get generic database object sql.DB to use its functions
	sqlDB, err := gormDB.DB()
	if err != nil {
		fmt.Println(err)

		return
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(20)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DBCube = gormDB
}

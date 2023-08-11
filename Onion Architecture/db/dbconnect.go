package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/pressly/goose"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func GetDB() *gorm.DB {
	return DB
}

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := dbUserName + ":" + dbPassword +
		"@(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database security_awareness...")
		fmt.Println(err)

		return
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		// singular table name
		// NamingStrategy: schema.NamingStrategy{
		// 	SingularTable: true,
		// },
	})
	if err != nil {
		fmt.Println("Failed to connect to database security_awareness (gorm)...")
		fmt.Println(err)

		return
	}

	// set db engine to "InnoDB"
	gormDB.Set("gorm:table_options", "ENGINE=InnoDB")

	// set charset to utf8mb4
	gormDB.Set("gorm:table_options", "CHARSET=utf8mb4")

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

	// Set the goose's dialect to mysql
	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	// // Runs the migrations
	// err = goose.Up(sqlDB, "db/migrations")
	// if err != nil {
	// 	log.Fatalf("failed to run migrations: %v", err)
	// }

	fmt.Println("12423534534534")

	DB = gormDB
}

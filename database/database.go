package database

import (
	"fmt"
	"log"
	"os"
	// "sync"
	"time"

	"github.com/hveda/todo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// Todo channels
var TODOID = make(chan int)
var CREATETODO = make(chan *models.Todo)
var UPDATETODO = make(chan *models.Todo)
var DELETETODO = make(chan *models.Todo)

// Activity channels
var ACTIVITYID = make(chan int)
var CREATEACTIVITY = make(chan *models.Activity)
var UPDATEACTIVITY = make(chan *models.Activity)
var DELETEACTIVITY = make(chan *models.Activity)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ConnectDb() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		getEnv("MYSQL_USER", "user"),
		getEnv("MYSQL_PASSWORD", "password"),
		getEnv("MYSQL_HOST", "localhost"),
		getEnv("MYSQL_PORT", "3306"),
		getEnv("MYSQL_DBNAME", "challenge_db"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:     true,
		CreateBatchSize: 1000,
		// SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(128)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(2 * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(1 * time.Minute))
	sqlDB.Ping()
	if err != nil {
		log.Fatal("Failed to config database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = Dbinstance{
		Db: db,
	}
}

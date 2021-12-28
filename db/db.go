package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"models"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func Init() {
		
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		getEnv("MYSQL_USER", "user"),
		getEnv("MYSQL_PASSWORD", "password"),
		getEnv("MYSQL_HOST", "localhost"),
		getEnv("MYSQL_PORT", "3306"),
		getEnv("MYSQL_DBNAME", "challenge_2_db"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
	if err != nil {
		panic(err.Error)
	}

	db.AutoMigrate(&models.ToDo{}, &models.ActivityGroup{})
	db.Set("gorm:table_options", "ENGINE=InnoDB")
}

func DbManager() *gorm.DB {
	return db
}

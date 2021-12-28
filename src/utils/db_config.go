package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func DbConnect() (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		getEnv("MYSQL_USER", "user"),
		getEnv("MYSQL_PASSWORD", "password"),
		getEnv("MYSQL_HOST", "localhost"),
		getEnv("MYSQL_PORT", "3306"),
		getEnv("MYSQL_DBNAME", "challenge_2_db"),
	)

	return gorm.Open(mysql.Open(url), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})
}
package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hveda/todo/src/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func load_db_config() (types.DbConnection, error) {
	db_config_file_data, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		return types.DbConnection{}, errors.New("db_config.json file not found")
	}

	var db_config types.DbConnection
	err = json.Unmarshal([]byte(db_config_file_data), &db_config)
	if err != nil {
		return types.DbConnection{}, err
	}
	return db_config, nil
}

func DbConnect() (*gorm.DB, error) {
	var url string
	config, err := load_db_config()
	if err != nil {
		// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			getEnv("MYSQL_USER", "user"),
			getEnv("MYSQL_PASSWORD", "password"),
			getEnv("MYSQL_HOST", "localhost"),
			getEnv("MYSQL_PORT", "3306"),
			getEnv("MYSQL_DBNAME", "challenge_2_db"),
		)
	} else {
		url = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", config.Username, config.Password, config.DbName)
	}
	return gorm.Open(mysql.Open(url), &gorm.Config{})
}
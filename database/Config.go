package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Config to maintain DB configuration properties
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var config = Config{
	ServerName: "34.121.36.6:3306",
	User:       "root",
	Password:   "P@ssw0rd",
	DB:         "mobility",
}

func InitConn() (db *sql.DB) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	db, err := sql.Open("mysql", connectionString)

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(version)
	return db
}

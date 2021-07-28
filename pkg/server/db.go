package server

import (
	"fmt"
	"os"
)

func NewDSN() string {
	user := os.Getenv("APP_DB_USER")
	password := os.Getenv("APP_DB_PASSWORD")
	host := os.Getenv("APP_DB_HOST")
	port := os.Getenv("APP_DB_PORT")
	db := os.Getenv("APP_DB_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, db)
}

package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"

	lg "github.com/bahodurnazarov/activity/pkg/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // add this
)

var DB *gorm.DB

func Conn() *sql.DB {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		lg.Errl.Fatal("Failed to connect to database. \n", err)
	}

	return db
}

package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var db *sql.DB

type connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Printf("Error loading .env file %s\n", err.Error())
		return
	}

	connInfo := connection{
		Host:     os.Getenv("POSTGRES_URL"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
	}

	db, err = sql.Open("postgres", connToString(connInfo))
	if err != nil {
		fmt.Printf("Error connecting to the Db %s\n", err.Error())
		return
	} else {
		fmt.Printf("Db connection open \n")
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error could not ping the Db %s\n", err.Error())
		return
	} else {
		fmt.Printf("Db ping successful \n")
	}
}

func connToString(info connection) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}

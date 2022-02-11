package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadDotEnv(env string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	val := os.Getenv(env)
	return val
}

func connectionDB() *sql.DB {
	Driver := loadDotEnv("DBDRIVER")
	User := loadDotEnv("DBUSER")
	Pass := loadDotEnv("DBPASS")
	Name := loadDotEnv("DBNAME")

	connection, err := sql.Open(Driver, User+":"+Pass+"@tcp(127.0.0.1)/"+Name)
	if err != nil {
		panic(err.Error())
	}
	return connection
}

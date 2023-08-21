package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// connect to db
	host := ""
	port := 5432 
	user := ""
	password := ""
	dbname := ""

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// db, err := sqlx.Connect("postgres", postgresqlDbInfo)

	database, err := gorm.Open(postgres.Open(postgresqlDbInfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

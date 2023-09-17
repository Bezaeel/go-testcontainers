package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbContext *gorm.DB

type DbOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func ConnectDatabase(opt DbOptions) {

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		opt.Host, opt.Port, opt.User, opt.Password, opt.Name)

	database, err := gorm.Open(postgres.Open(postgresqlDbInfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DbContext = database
}

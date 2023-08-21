package utils

import (
	"fmt"
	"go-testcontainers/customer"
	"go-testcontainers/database"
)

func Migrate() {
	database.DB.AutoMigrate(&customer.Customer{})
	fmt.Println("👍 Migration complete")

}

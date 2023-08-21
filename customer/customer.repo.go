package customer

import (
	"fmt"
	"go-testcontainers/database"
)

type CustomerRepo interface {
	Create(customer *Customer) *Customer
}

type customerRepo struct{}

func NewCustomerRepo() CustomerRepo {
	return &customerRepo{}
}

// Create implements CustomerRepo
func (cr *customerRepo) Create(customer *Customer) *Customer {
	database.DB.Create(&customer)
	fmt.Print(customer);
	return customer
}

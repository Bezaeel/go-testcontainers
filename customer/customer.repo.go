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
	database.DbContext.Create(&customer)
	fmt.Print(customer)
	return customer
}

func (cr *customerRepo) Detail(id int) *Customer {
	var customer Customer
	database.DbContext.First(&customer, 1)
	return &customer
}

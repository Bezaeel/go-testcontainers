package customer

import (
	"go-testcontainers/database"
)

type CustomerService interface {
	CreateCustomer(customer Customer) (*Customer, error)
	Details(id int) (*Customer, error)
}

type customerService struct {
}

func NewCustomerService() CustomerService {
	return &customerService{}
}

// CreateCustomer implements CustomerService
func (cs *customerService) CreateCustomer(customer Customer) (*Customer, error) {
	err := database.DbContext.Create(&customer)
	if err.Error != nil {
		return nil, err.Error
	}
	return &customer, nil
}

// CreateCustomer implements CustomerService
func (cs *customerService) Details(id int) (*Customer, error) {
	var result Customer
	err := database.DbContext.First(&result, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &result, nil
}

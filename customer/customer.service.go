package customer

import (
	"fmt"
	"sync"
)


type CustomerService interface {
	CreateCustomer(customer Customer) *Customer
	AskTalabiAsync(chanResult chan CustomerResult, num int)
	AskTalabiAsyncBackground(wg *sync.WaitGroup, num int)
}

type customerService struct {
	customerRepo CustomerRepo
}

type CustomerResult struct {
	Err error `json:"err"`
	Sum int   `json:"sum"`
}

func NewCustomerService() CustomerService {
	return &customerService{
		customerRepo: NewCustomerRepo(),
	}
}

// CreateCustomer implements CustomerService
func (cs *customerService) CreateCustomer(customer Customer) *Customer {
	result := cs.customerRepo.Create(&customer)
	return result
}

// CreateCustomer implements CustomerServiceAsync
func (cs *customerService) AskTalabiAsync(chanResult chan CustomerResult, num int) {
	result := CustomerResult{
		Err: nil,
		Sum: num * 2,
	}
	chanResult <- result
	close(chanResult)
}

// CreateCustomer implements CustomerServiceAsync
func (cs *customerService) AskTalabiAsyncBackground(wg *sync.WaitGroup, num int) {
	wg.Add(1)
	defer wg.Done()

	fmt.Printf("sum::::%v", num*2)
}

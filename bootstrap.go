package main

import "go-testcontainers/customer"

func Init() {
	customer.NewCustomerController(routes)
}

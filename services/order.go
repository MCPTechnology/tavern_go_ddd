// Package services holds all the services that connects repositories into a business flow
package services

import (
	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/domain/customer"
	"github.com/matheuscaputopires/ddd-go/domain/customer/memory"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type OrderConfiguration func(os *OrderService) error

// OrderService is a implementation of the OrderService
type OrderService struct {
	customer customer.CustomerRepository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the order service
	os := &OrderService{}
	// Apply all configurations passed
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// This is done so that the parent function can take in all the needed parameters
	return func (os *OrderService) error {
		os.customer = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if paramaeters such as connection strings, they could be input here
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	// Get the customer
	c, err := o.customer.Get(customerID)
	if err != nil {
		return err
	}
	
	return nil
}

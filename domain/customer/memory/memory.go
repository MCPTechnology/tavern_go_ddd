// Package memory is a in-memory implementation of the customer repository
package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/tavern/domain/customer"
)

// MemoryRepository fulfills the CustomerRepository interface
type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

// New is a function to generate a new repository of customers
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

// Get finds a customer by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFount
}

// Add will add a new customer to the repository
func (mr *MemoryRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}
	// Make sure the customer isn't already in the repo
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("The customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	// If it doesn't exist, we add it
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryRepository) Update(c customer.Customer) error {
	// Make sure the customer exists
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("The customer doesn't exist: %w", customer.ErrUpdateCustomer)
	}
	// If it doesn't exist, we add it
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

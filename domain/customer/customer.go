// Package customer
// File: customer.go
// Customer is an aggregate that represents a customer
package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/tavern"
)

var ErrInvalidPerson = errors.New("A customer should have a valid name")

type Customer struct {
	// Root Customer Entity, that contains a ID (Person), products and transactions
	// Fields are lowercase (not acessible directly to other domains) and the aggregate does not decide
	// what structure each data should have
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate according to the factory design pattern
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
		Age:  0,
	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil
}

// GetID gets the root ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}

// GetName returns the Customer name
func (c *Customer) GetName() string {
	return c.person.Name
}

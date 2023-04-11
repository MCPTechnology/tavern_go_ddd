// Package customer
// File: repository.go
// Package Customer holds all the domain logic for the customer domain
package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFount = errors.New("The cusrtomer was not found in the repository.")
	ErrFailedToAddCustomer = errors.New("Unable to add the customer to the repository.")
	ErrUpdateCustomer = errors.New("Unable to update the customer in the repository.")
)

// Repository is an interface that defines the rules arount what a cusromer repository
// has to be able to perform
type Repository interface{
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	Update(Customer) error
}

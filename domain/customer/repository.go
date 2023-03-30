// Package Customer holds all the domain logic for the customer domain
package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/aggregate"
)

var (
	ErrCustomerNotFount = errors.New("The cusrtomer was not found in the repository.")
	ErrFailedToAddCustomer = errors.New("Unable to add the customer to the repository.")
	ErrUpdateCustomer = errors.New("Unable to update the customer in the repository.")
)

// CustomerRepository is an interface that defines the rules arount what a cusromer repository
// has to be able to perform
type CustomerRepository interface{
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}

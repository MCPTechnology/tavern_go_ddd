// Package product
// It holds the repository implementation and the implementation for a ProductRepository
package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/aggregate"
)

var (
	ErrProductNotFound = errors.New("The product was not found")
	ErrProductAlreadyExists = errors.New("The product already exists")
)

type ProductRepository interface{
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}

// Package memory is an in memory implementation of the ProductRepository interface
package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/aggregate"
	"github.com/matheuscaputopires/ddd-go/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

// GetAll returns all products as a slice
// This will not retuns an error but a database implementation might
func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	// Collects all products from map
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

// GetByID searches for a product based on it's ID
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

// Add will Add a new product to the repository
func (mpr *MemoryProductRepository) Add(newProduct aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[newProduct.GetID()] = newProduct
	return nil
}

// Update will change all values for a product based on it's ID
func (mpr *MemoryProductRepository) Update(upprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// Delete remove an product from the repository
func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
// Package aggregate
// File: product.go
// Product is an aggregate that represents a product
package aggregate

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/entity"
)

// ErrMissingValues is returned when a product is created without a name or description
var ErrMissingValues = errors.New("Missing values")

// Product is an aggregate that combines item with a price and quantity
type Product struct {
	// Item is the root entity
	item  *entity.Item
	price float64
	// Quantity is the number of products in stock
	quantity int
}

// New product will return an error if name of description is empty
func NewProduct(name string, description string, price float64) (Product, error) {
	if strings.TrimSpace(name) == "" || strings.TrimSpace(description) == "" {
		return Product{}, ErrMissingValues
	}
	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Desctiption: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}

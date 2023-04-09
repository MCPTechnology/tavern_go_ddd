package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/aggregate"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peanuts, err := aggregate.NewProduct("Peanuts", "Snack", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Drink", 5.99)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{beer, peanuts, wine}
	return products
}

func TestOrder_NewOrderService(t *testing.T) {
	// Create a few products to insert into in memory repo
	products := init_products(t)

	// Create order service and add products to inventory
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	// Add a Customer
	cust, err := aggregate.NewCustomer("TestCustomer")
	if err != nil {
		t.Error(err)
	}

	err = os.customer.Add(cust)
	if err != nil {
		t.Error(err)
	}

	// Request an Order for one beer
	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}

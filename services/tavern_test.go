package services

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/ddd-go/aggregate"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("TestCustomer")
	if err != nil {
		t.Error(err)
	}

	err = os.customer.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}

func TestTavern_Mongo(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := NewOrderService(
		WithMongoCustomerRepository(context.Background(), "mongodb://admin:admin123@mongo:27017/"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("TestCustomer")
	if err != nil {
		t.Error(err)
	}

	err = os.customer.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}

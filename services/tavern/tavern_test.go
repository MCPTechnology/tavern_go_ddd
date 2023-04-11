package tavern

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/tavern/domain/product"
	"github.com/matheuscaputopires/tavern/services/order"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peanuts, err := product.NewProduct("Peanuts", "Snack", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Drink", 5.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{beer, peanuts, wine}
	return products
}


func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("TestCustomer")
	if err != nil {
		t.Fatal(err)
	}
	
	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}

func TestTavern_Mongo(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://admin:admin123@mongo:27017/"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("TestCustomer")
	if err != nil {
		t.Fatal(err)
	}
	
	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}

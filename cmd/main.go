package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/tavern/domain/product"
	"github.com/matheuscaputopires/tavern/services/order"
	"github.com/matheuscaputopires/tavern/services/tavern"
)

func main() {
	products := productInventory()

	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(context.Background(), "mongodb://admin:admin123@localhost:27017/tavern"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tavern, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("RealCustomer")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peanuts, err := product.NewProduct("Peanuts", "Snack", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Drink", 5.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{beer, peanuts, wine}
	return products
}

// Package services holds all the services that connects repositories into a business flow
package order

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/matheuscaputopires/tavern/domain/customer"
	customerMemoryRepository "github.com/matheuscaputopires/tavern/domain/customer/memory"
	customerMongoDBRepository "github.com/matheuscaputopires/tavern/domain/customer/mongodb"
	"github.com/matheuscaputopires/tavern/domain/product"
	productDomain "github.com/matheuscaputopires/tavern/domain/product"
	productMemoryRepository "github.com/matheuscaputopires/tavern/domain/product/memory"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type OrderConfiguration func(os *OrderService) error

// OrderService is a implementation of the OrderService
type OrderService struct {
	customer customer.Repository
	products productDomain.Repository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the order service
	os := &OrderService{}
	// Apply all configurations passed
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(customer customer.Repository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// This is done so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customer = customer
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if paramaeters such as connection strings, they could be input here
	customer := customerMemoryRepository.New()
	return WithCustomerRepository(customer)
}

// WithMongoCustomerRepository applies a MongoDB customer repository to the OrderService
func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the mongo repo
		cr, err := customerMongoDBRepository.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customer = cr
		return nil
	}
}

// WithMemoryProductRepository adds an in memory product repo and adds all input products to it
func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if parameters are needed such as connection strings, this is where it should be
		productRepo := productMemoryRepository.New()

		// Add Items to repo
		for _, p := range products {
			err := productRepo.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = productRepo
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	customer, err := o.customer.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []product.Product
	var price float64
	for _, productID := range productIDs {
		product, err := o.products.GetByID(productID)
		if err != nil {
			return 0, err
		}

		products = append(products, product)
		price += product.GetPrice()
	}

	// All products are in store, so we can create the order
	log.Printf("Customer %s has ordered %d products", customer.GetID(), len(products))

	return price, nil
}

func (os *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	customer, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	
	err = os.customer.Add(customer)
	if err != nil {
		return uuid.Nil, err
	}

	return customer.GetID(), nil
}
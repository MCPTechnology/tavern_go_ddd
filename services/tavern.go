package services

import (
	"log"

	"github.com/google/uuid"
)

// TavernConfiguration is an alias that takes a pointer and modifies the Tavern
type TavernConfiguration func(ts *TavernService) error

type TavernService struct {
	// OrderService is used to handle orders
	OrderService *OrderService
	// BillingService is used to handle billing
	// TODO needs implementation
	BillingService interface{}
}

// NewTavern takes a variable amount of TavernConfigurations
func NewTavern(cfgs ...TavernConfiguration) (*TavernService, error) {
	// Create the TavernService
	ts := &TavernService{}

	// Apply all configurations passed to it
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(ts)
		if err != nil {
			return nil, err
		}

	}
	return ts, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	// Returns a function that matches the TavernConfiguration signature
	return func(ts *TavernService) error {
		ts.OrderService = os
		return nil
	}
}

// Order performs an order for a customer
func (ts *TavernService) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := ts.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	// Actually Bill the customer
	log.Printf("Bill the Customer %v: %0.0f", customer, price)
	// err = ts.BillingService.Bill(customer, price)
	return nil
}

package taveran

import (
	"github.com/MohammadAlhallaq/phpToGo/services/order"
	"github.com/google/uuid"
	"log"
)

type TavernConfiguration func(os *TavernService) error

type TavernService struct {
	OrderService   *order.OrderService
	BillingService interface{}
}

func NewTaveran(cfgs ...TavernConfiguration) (*TavernService, error) {
	t := &TavernService{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func withOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *TavernService) error {
		t.OrderService = os
		return nil
	}
}

func (t *TavernService) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the Customer: %0.0f", price)

	//add the billing magic here

	return nil
}

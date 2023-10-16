package services

import (
	"github.com/google/uuid"
	"phpToGo/aggregate"
	"testing"
)

func init_product(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []aggregate.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrderService_CreateOrder(t *testing.T) {

	products := init_product(t)
	os, err := NewOrderService(WithMemoryCustomerRepository(), WithMemoryProductRepository(products))

	if err != nil {
		t.Error(err)
	}
	customer, err := aggregate.NewCustomer("mohammad")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(customer)
	if err != nil {
		t.Error(err)
	}

	var orders []uuid.UUID
	for _, product := range products {
		orders = append(orders, product.GetID())
	}

	_, err = os.CreateOrder(customer.GetID(), orders)
	if err != nil {
		t.Error(err)
	}
}

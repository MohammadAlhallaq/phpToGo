package services

import (
	"github.com/google/uuid"
	"phpToGo/aggregate"
	"testing"
)

func TestNewTaveran(t *testing.T) {

	products := init_product(t)
	os, err := NewOrderService(WithMemoryProductRepository(products), WithMongoCustomerRepository("mongodb://127.0.0.1:27017"))
	if err != nil {
		t.Error(err)
	}

	taverna, err := NewTaveran(withOrderService(os))
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

	err = taverna.Order(customer.GetID(), orders)
	if err != nil {
		t.Error(err)
	}
}

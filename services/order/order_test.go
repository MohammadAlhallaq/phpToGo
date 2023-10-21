package order

import (
	"github.com/MohammadAlhallaq/phpToGo/domain/customer"
	"github.com/MohammadAlhallaq/phpToGo/domain/product"
	"github.com/google/uuid"
	"testing"
)

func initProduct(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestOrderService_CreateOrder(t *testing.T) {

	products := initProduct(t)
	os, err := NewOrderService(WithMemoryCustomerRepository(), WithMemoryProductRepository(products))
	if err != nil {
		t.Error(err)
	}
	customer, err := customer.NewCustomer("mohammad")
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

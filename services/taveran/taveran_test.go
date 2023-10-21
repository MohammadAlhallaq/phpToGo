package taveran

import (
	"github.com/MohammadAlhallaq/phpToGo/domain/product"
	"github.com/MohammadAlhallaq/phpToGo/services/order"
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

func TestNewTaveran(t *testing.T) {

	products := initProduct(t)
	os, err := order.NewOrderService(order.WithMemoryProductRepository(products), order.WithMongoCustomerRepository("mongodb://127.0.0.1:27017"))
	if err != nil {
		t.Error(err)
	}

	taverna, err := NewTaveran(withOrderService(os))
	if err != nil {
		t.Error(err)
	}

	userID, err := os.AddCustomer("mohammad")
	if err != nil {
		t.Error(err)
	}

	var orders []uuid.UUID
	for _, product := range products {
		orders = append(orders, product.GetID())
	}

	err = taverna.Order(userID, orders)
	if err != nil {
		t.Error(err)
	}
}

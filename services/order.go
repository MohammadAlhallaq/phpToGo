package services

import (
	"github.com/google/uuid"
	"log"
	"phpToGo/aggregate"
	"phpToGo/domain/customer"
	memoryCustomer "phpToGo/domain/customer/memory"
	"phpToGo/domain/product"
	memoryProduct "phpToGo/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepo
	products  product.ProductRepo
}

func NewOrderService(configurations ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range configurations {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memoryCustomer.New()

	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := memoryProduct.New()

		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
}

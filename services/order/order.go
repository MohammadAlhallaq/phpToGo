package order

import (
	"context"
	"github.com/MohammadAlhallaq/phpToGo/domain/customer"
	memoryCustomer "github.com/MohammadAlhallaq/phpToGo/domain/customer/memory"
	"github.com/MohammadAlhallaq/phpToGo/domain/customer/mongo"
	"github.com/MohammadAlhallaq/phpToGo/domain/product"
	memoryProduct "github.com/MohammadAlhallaq/phpToGo/domain/product/memory"
	"github.com/google/uuid"
	"log"
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

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
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

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []product.Product
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

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}

package services

import (
	"phpToGo/domain/customer"
	"phpToGo/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepo
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

func WithCustomerRepository(cr customer.CustomerRepo) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}
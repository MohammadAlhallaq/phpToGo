package product

import (
	"errors"
	"github.com/MohammadAlhallaq/phpToGo"
	"github.com/google/uuid"
)

var (
	ErrMissingValue = errors.New("missing value")
)

type Product struct {
	item     *phpToGo.Item
	price    float64
	quantity int
}

func NewProduct(name string, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	return Product{
		item: &phpToGo.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *phpToGo.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}

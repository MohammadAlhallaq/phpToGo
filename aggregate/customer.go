package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"phpToGo/entity"
	"phpToGo/valueobject"
)

var ErrInvalidCustomer = errors.New("name is required")

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {

	if name == "" {
		return Customer{}, ErrInvalidCustomer
	}

	return Customer{
		person: &entity.Person{
			Name: name,
			ID:   uuid.New(),
		},
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}
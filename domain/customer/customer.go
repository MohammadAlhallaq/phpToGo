package customer

import (
	"errors"
	"github.com/MohammadAlhallaq/phpToGo"
	"github.com/google/uuid"
)

var ErrInvalidCustomer = errors.New("name is required")

type Customer struct {
	person       *phpToGo.Person
	products     []*phpToGo.Item
	transactions []phpToGo.Transaction
}

func NewCustomer(name string) (Customer, error) {

	if name == "" {
		return Customer{}, ErrInvalidCustomer
	}

	return Customer{
		person: &phpToGo.Person{
			Name: name,
			ID:   uuid.New(),
		},
		products:     make([]*phpToGo.Item, 0),
		transactions: make([]phpToGo.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &phpToGo.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &phpToGo.Person{}
	}
	c.person.Name = name
}

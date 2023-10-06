package memory

import (
	"fmt"
	"github.com/google/uuid"
	"phpToGo/aggregate"
	"phpToGo/domain/customer"
	"sync"
)

type MemoryRepo struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepo {
	return &MemoryRepo{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (m *MemoryRepo) Get(id uuid.UUID) (aggregate.Customer, error) {
	if c, ok := m.customers[id]; ok {
		return c, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (m *MemoryRepo) Add(c aggregate.Customer) error {
	if m.customers == nil {
		m.Lock()
		m.customers = make(map[uuid.UUID]aggregate.Customer)
		m.Unlock()
	}
	if _, ok := m.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil
}

func (m *MemoryRepo) Update(c aggregate.Customer) error {

	if _, ok := m.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil
}

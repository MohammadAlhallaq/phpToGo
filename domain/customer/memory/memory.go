package memory

import (
	"fmt"
	"github.com/MohammadAlhallaq/phpToGo/domain/customer"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepo struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepo {
	return &MemoryRepo{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (m *MemoryRepo) Get(id uuid.UUID) (customer.Customer, error) {
	if c, ok := m.customers[id]; ok {
		return c, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (m *MemoryRepo) Add(c customer.Customer) error {
	m.Lock()
	defer m.Unlock()

	if m.customers == nil {
		m.customers = make(map[uuid.UUID]customer.Customer)
	}
	if _, ok := m.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	m.customers[c.GetID()] = c
	return nil
}

func (m *MemoryRepo) Update(c customer.Customer) error {

	if _, ok := m.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	m.Lock()
	defer m.Unlock()
	m.customers[c.GetID()] = c
	return nil
}

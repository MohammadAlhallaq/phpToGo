package memory

import (
	"github.com/google/uuid"
	"phpToGo/aggregate"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (m *MemoryRepository) Get(uuid uuid.UUID) {
	panic("implement me")
}

func (m *MemoryRepository) Add(customer aggregate.Customer) {
	panic("implement me")
}

func (m *MemoryRepository) Update(customer aggregate.Customer) {
	panic("implement me")
}

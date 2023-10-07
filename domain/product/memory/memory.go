package memory

import (
	"github.com/google/uuid"
	"phpToGo/aggregate"
	"phpToGo/domain/product"
	"sync"
)

type MemoryProductRepo struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepo {
	return &MemoryProductRepo{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (m *MemoryProductRepo) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range m.products {
		products = append(products, product)
	}

	return products, nil
}

func (m *MemoryProductRepo) GetByID(id uuid.UUID) (aggregate.Product, error) {

	if p, ok := m.products[id]; ok {
		return p, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (m *MemoryProductRepo) Add(newProduct aggregate.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	m.products[newProduct.GetID()] = newProduct
	return nil
}

func (m *MemoryProductRepo) Update(updProduct aggregate.Product) error {
	m.Lock()
	m.Unlock()

	if _, ok := m.products[updProduct.GetID()]; !ok {
		return product.ErrProductNotFound
	}
	m.products[updProduct.GetID()] = updProduct
	return nil
}

func (m *MemoryProductRepo) Delete(id uuid.UUID) error {
	m.Lock()
	m.Unlock()

	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(m.products, id)
	return nil
}

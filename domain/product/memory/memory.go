package memory

import (
	"github.com/MohammadAlhallaq/phpToGo/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepo struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryRepo {
	return &MemoryRepo{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (m *MemoryRepo) GetAll() []product.Product {
	var products []product.Product

	for _, prod := range m.products {
		products = append(products, prod)
	}

	return products
}

func (m *MemoryRepo) GetByID(id uuid.UUID) (product.Product, error) {

	if p, ok := m.products[id]; ok {
		return p, nil
	}

	return product.Product{}, product.ErrProductNotFound
}

func (m *MemoryRepo) Add(newProduct product.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[newProduct.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	m.products[newProduct.GetID()] = newProduct
	return nil
}

func (m *MemoryRepo) Update(updProduct product.Product) error {
	m.Lock()
	m.Unlock()

	if _, ok := m.products[updProduct.GetID()]; !ok {
		return product.ErrProductNotFound
	}
	m.products[updProduct.GetID()] = updProduct
	return nil
}

func (m *MemoryRepo) Delete(id uuid.UUID) error {
	m.Lock()
	m.Unlock()

	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(m.products, id)
	return nil
}

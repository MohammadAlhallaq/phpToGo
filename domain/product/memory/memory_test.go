package memory

import (
	"errors"
	"github.com/google/uuid"
	"phpToGo/aggregate"
	"phpToGo/domain/product"
	"testing"
)

func TestMemoryProductRepo_Add(t *testing.T) {
	type testCase struct {
		test        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "adding product",
			expectedErr: nil,
		},
		{
			test:        "adding product",
			expectedErr: product.ErrProductAlreadyExist,
		},
	}

	repo := New()
	newProd, err := aggregate.NewProduct("demo", "Good for you're health", 1.99)
	if err != nil {
		t.Error(err)
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := repo.Add(newProd)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryProductRepo_GetAll(t *testing.T) {
	repo := New()
	newProd1, err := aggregate.NewProduct("demo1", "Good for you're health1", 1.99)
	newProd2, err := aggregate.NewProduct("demo2", "Good for you're health2", 1.99)
	err = repo.Add(newProd2)
	err = repo.Add(newProd1)
	if err != nil {
		t.Error(err)
	}
	allProducts := repo.GetAll()

	if len(allProducts) != 2 {
		t.Error(err)
	}
}

func TestMemoryProductRepo_GetByID(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	repo := New()
	newProd1, err := aggregate.NewProduct("demo1", "Good for you're health1", 1.99)
	err = repo.Add(newProd1)
	if err != nil {
		t.Error(err)
	}

	testCases := []testCase{
		{
			test:        "Get By ID",
			id:          newProd1.GetID(),
			expectedErr: nil,
		},
		{
			test:        "Product not found",
			id:          uuid.New(),
			expectedErr: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

package product_test

import (
	"errors"
	"github.com/MohammadAlhallaq/phpToGo/domain/product"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			expectedErr: product.ErrMissingValue,
		},
		{
			test:        "validvalues",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := product.NewProduct(tc.name, tc.description, tc.price)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

func TestProduct_GetID(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}
	tc := testCase{
		test:        "should return product",
		name:        "test",
		description: "test",
		price:       1.0,
		expectedErr: nil,
	}
	t.Run(tc.test, func(t *testing.T) {
		product, err := product.NewProduct(tc.name, tc.description, tc.price)
		if !errors.Is(err, tc.expectedErr) {
			t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
		}

		id := product.GetID()
		if reflect.TypeOf(id) != reflect.TypeOf(uuid.UUID{}) {
			t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
		}
	})
}

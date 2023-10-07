package memory

import (
	"errors"
	"github.com/google/uuid"
	"phpToGo/aggregate"
	"phpToGo/domain/customer"
	"reflect"
	"testing"
)

func TestMemoryRepository_Get(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("mohammad")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	repo := MemoryCustomerRepo{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			test:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			test:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemoryRepository_Add(t *testing.T) {
	type testCase struct {
		test        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Add Customer",
			cust:        "Percy",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			repo := MemoryCustomerRepo{
				customers: map[uuid.UUID]aggregate.Customer{},
			}
			cust, err := aggregate.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}
			err = repo.Add(cust)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}

func TestMemoryRepository_Update(t *testing.T) {
	type testCase struct {
		test        string
		cust        string
		updatedCust string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Update Customer",
			cust:        "Percy",
			updatedCust: "Mohammad",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			repo := MemoryCustomerRepo{
				customers: map[uuid.UUID]aggregate.Customer{},
			}
			cust, err := aggregate.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			cust.SetName(tc.updatedCust)
			err = repo.Update(cust)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())

			if err != nil {
				t.Fatal(err)
			}
			if found.GetName() == tc.cust {
				t.Errorf("Expected %v, got %v", found.GetName(), tc.cust)
			}
		})
	}
}

func TestNew(t *testing.T) {

	type testCase struct {
		test        string
		expectedErr error
	}
	testCases := testCase{
		test:        "New Customer",
		expectedErr: nil,
	}

	t.Run(testCases.test, func(t *testing.T) {
		CustomerAggregate := New()
		repo := &MemoryCustomerRepo{
			customers: map[uuid.UUID]aggregate.Customer{},
		}
		if reflect.TypeOf(CustomerAggregate) != reflect.TypeOf(repo) {
			t.Error("customerAggregate is not of type Customer")
		}
	})
}

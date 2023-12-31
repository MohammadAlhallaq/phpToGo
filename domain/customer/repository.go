package customer

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

type CustomerRepo interface {
	Get(id uuid.UUID) (Customer, error)
	Add(customer Customer) error
	Update(customer Customer) error
}

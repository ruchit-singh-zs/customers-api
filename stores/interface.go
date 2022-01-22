package stores

import "customerApi/models"

type Customer interface {
	CreateCustomer(c models.Customer) error
	GetCustomer(id string) (models.Customer, error)
	UpdateCustomer(id string, c models.Customer) error
	DeleteCustomer(id string) error
}

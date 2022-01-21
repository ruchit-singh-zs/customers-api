package services

import "customerApi/models"

type Customer interface {
	CreateCustomer(customer models.Customer) error
	GetCustomer(id string) error
	UpdateCustomer(id string) error
	DeleteCustomer(id string) error
}

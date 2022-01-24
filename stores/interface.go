package stores

import "customerApi/models"

type Customer interface {
	Create(c models.Customer) error
	Get(id int) (models.Customer, error)
	Update(id int, c models.Customer) error
	Delete(id int) error
}

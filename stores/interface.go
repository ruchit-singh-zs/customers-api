package stores

import (
	"customers-api/models"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type Customer interface {
	Create(ctx *gofr.Context, c models.Customer) (models.Customer, error)
	GetByID(ctx *gofr.Context, id string) (c models.Customer, e error)
	UpdateByID(ctx *gofr.Context, id string, c models.Customer) error
	DeleteByID(ctx *gofr.Context, id string) (models.Customer, error)
}

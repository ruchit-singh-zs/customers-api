package customer

import (
	"customers-api/models"
	"customers-api/stores"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"log"
)

type store struct {
}

func New() stores.Customer {
	return store{}
}

func (s store) Create(ctx *gofr.Context, customer models.Customer) (models.Customer, error) {
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customer (id,name,phoneNo,address) VALUES (?,?, ?, ?)",
		customer.ID, customer.Name, customer.PhoneNo, customer.Address)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (s store) GetByID(ctx *gofr.Context, id string) (models.Customer, error) {
	var c models.Customer
	err := ctx.DB().QueryRowContext(ctx, "SELECT * FROM customer WHERE id = ?", id).Scan(&c.ID, &c.Name, &c.PhoneNo, &c.Address)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (s store) UpdateByID(ctx *gofr.Context, id string, c models.Customer) error {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE customer SET name = ?, phoneNo = ?, address = ? WHERE id = ?",
		c.Name, c.PhoneNo, c.Address, id)

	log.Println(c.ID, c.Name, c.Address)

	if err != nil {
		return err
	}

	return nil
}

func (s store) DeleteByID(ctx *gofr.Context, id string) (models.Customer, error) {
	var c models.Customer
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM customer WHERE id = ?", id)
	if err != nil {
		return c, err
	}
	return c, nil
}

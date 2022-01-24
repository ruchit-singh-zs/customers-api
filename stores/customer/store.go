package customer

import (
	"customerApi/errors"
	"database/sql"

	"customerApi/models"
	"customerApi/stores"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) stores.Customer {
	return store{db: db}
}

func (s store) Create(customer models.Customer) error {
	_, err := s.db.Exec("INSERT INTO Customer (ID,NAME , PHONENO, ADDRESS) VALUES (?,?, ?, ?)",
		customer.ID, customer.Name, customer.PhoneNo, customer.Address)

	return err
}

func (s store) Get(id int) (models.Customer, error) {
	var c models.Customer
	err := s.db.QueryRow("SELECT * FROM Customer WHERE ID = ?", id).
		Scan(&c.ID, &c.Name, &c.PhoneNo, &c.Address)

	switch err.(type) {
	case errors.NoEntity:
		return c, err
	case nil:
		return c, nil
	default:
		return c, err
	}
}

func (s store) Update(id int, customer models.Customer) error {
	_, err := s.db.Exec("UPDATE Customer SET NAME = ?, PHONENO=?, ADDRESS=? WHERE ID = ?",
		&customer.Name, &customer.PhoneNo, &customer.Address, id)
	return err
}

func (s store) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM Customer WHERE ID =?", id)
	return err
}

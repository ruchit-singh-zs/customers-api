package customer

import (
	"customerApi/models"
	"customerApi/stores"
	"database/sql"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) stores.Customer {
	return store{db: db}
}

func (s store) CreateCustomer(c models.Customer) error {
	_, err := s.db.Exec("INSERT INTO Customer (ID,NAME , PHONENO, ADDRESS) VALUES (?,?, ?, ?)",
		c.ID, c.Name, c.PhoneNo, c.Address)

	return err
}

func (s store) GetCustomer(id string) (models.Customer, string) {
	var c models.Customer
	err := s.db.QueryRow("SELECT * FROM Customer WHERE ID = ?", id).
		Scan(&c.ID, &c.Name, &c.PhoneNo, &c.Address)

	switch err {
	case sql.ErrNoRows:
		return c, "No record"
	case nil:
		return c, "No error"
	default:
		return c, "Internal Server error"
	}
}

func (s store) UpdateCustomer(id string, c models.Customer) error {
	_, err := s.db.Exec("UPDATE Customer SET NAME = ?, PHONENO=?, ADDRESS=? WHERE ID = ?",
		&c.Name, &c.PhoneNo, &c.Address, id)
	return err
}

func (s store) DeleteCustomer(id string) error {
	_, err := s.db.Exec("DELETE FROM Customer WHERE ID =?", id)
	return err
}

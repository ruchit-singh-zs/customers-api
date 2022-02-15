package models

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	PhoneNo string `json:"phoneNo"`
	Address string `json:"address"`
}

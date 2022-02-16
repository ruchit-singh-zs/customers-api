package main

import (
	handlers "customers-api/handlers/customer"
	stores "customers-api/stores/customer"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()
	s := stores.New()
	h := handlers.New(s)
	app.Server.ValidateHeaders = false
	app.GET("/customer/{id}", h.GetByID)
	app.POST("/customer", h.Create)
	app.PUT("/customer/update/{id}", h.UpdateByID)
	app.DELETE("/customer/delete/{id}", h.DeleteByID)
	app.Start()
}

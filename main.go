package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"customerApi/drivers"
	handlers "customerApi/handlers/customer"
	"customerApi/middleware"
	stores "customerApi/stores/customer"
)

func main() {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		log.Fatalf("FATAL, Can't Connect to database")
	}
	defer db.Close()
	s := stores.New(db)
	h := handlers.New(s)

	r := mux.NewRouter()
	r.HandleFunc("/customer/{id}", h.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customer", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/delete/{id}", h.DeleteByID).Methods(http.MethodDelete)
	r.HandleFunc("/customer/update/{id}", h.UpdateByID).Methods(http.MethodPut)

	r.Use(middleware.SetContentType)

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Cant Connect!")
	}
}

package main

import (
	handlers "customers-api/handlers/customer"
	stores "customers-api/stores/customer"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

//func main() {
//	db, err := drivers.ConnectToSQL()
//	if err != nil {
//		log.Fatalf("FATAL, Can't Connect to database")
//	}
//	defer db.Close()
//	s := stores.New(db)
//	h := handlers.New(s)
//
//	r := mux.NewRouter()
//	r.HandleFunc("/customer/{id}", h.GetByID).Methods(http.MethodGet)
//	r.HandleFunc("/customer", h.Create).Methods(http.MethodPost)
//	r.HandleFunc("/customer/delete/{id}", h.DeleteByID).Methods(http.MethodDelete)
//	r.HandleFunc("/customer/update/{id}", h.UpdateByID).Methods(http.MethodPut)
//
//	r.Use(middleware.SetContentType)
//
//	err = http.ListenAndServe(":8080", r)
//	if err != nil {
//		log.Println("Cant Connect!")
//	}
//}

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

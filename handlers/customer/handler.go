package customer

import (
	"customerApi/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"

	"customerApi/models"
)

type handler struct {
	service services.Customer
}

func New(s services.Customer) handler {
	return handler{service: s}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var c models.Customer
	err = json.Unmarshal(body, &c)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.service.CreateCustomer(c)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, errs := w.Write([]byte("Error in Inserting"))

		if errs != nil {
			log.Println(errs)
		}

		return
	}

	_, err = w.Write([]byte("successfully created"))
	if err != nil {
		log.Println(err)
	}
}

func (h handler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var c models.Customer

	err := h.service.GetCustomer(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte("No Record Exists"))

		if err != nil {
			log.Println(err)
		}

		resp, err := json.Marshal(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(resp)

		if err != nil {
			log.Println(err)
		}
	}
}

func (h handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var c models.Customer
	err = json.Unmarshal(body, &c)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UpdateCustomer(id)
	if err != nil {
		log.Printf("Error in Updating: %v", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = w.Write([]byte("Updated Successfully"))
	if err != nil {
		log.Println(err)
	}
}

func (h handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := h.service.DeleteCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error in deleting", err)

		return
	}

	_, err = w.Write([]byte("Deleted Successfully"))
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		log.Println(err)
	}
}

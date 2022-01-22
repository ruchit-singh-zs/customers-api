package customer

import (
	"customerApi/stores"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"

	"customerApi/drivers"
	"customerApi/models"
)

type handler struct {
	store stores.Customer
}

func New(store stores.Customer) handler {
	return handler{store: store}
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

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

	err = h.store.CreateCustomer(c)
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
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	c, err := h.store.GetCustomer(id)
	switch err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte("No Record Exists"))

		if err != nil {
			log.Println(err)
		}
	case nil:
		resp, err := json.Marshal(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(resp)

		if err != nil {
			log.Println(err)
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

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

	err = h.store.UpdateCustomer(id, c)
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
	db, err := drivers.ConnectToSQL()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	err = h.store.DeleteCustomer(id)
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

package customer

import (
	"bytes"
	"customerApi/models"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	id int
}

func (m mockService) CreateCustomer(customer models.Customer) error {
	if m.id == 2 {
		return errors.New("Server Error")
	}
	return nil
}
func (m mockService) GetCustomer(id string) error {
	if m.id == 2 {
		return errors.New("entity doesn't exist")
	}
	return nil
}
func (m mockService) UpdateCustomer(id string) error {
	if m.id == 2 {
		return errors.New("entity doesn't exist")
	}
	return nil
}
func (m mockService) DeleteCustomer(id string) error {
	if m.id == 2 {
		return errors.New("entity doesn't exist")
	}
	return nil
}

func TestCreate(t *testing.T) {
	res := `{"id":7,"name":"Shreya S","phoneNo":"9909111122","address":"BG Road Bangalore"}`
	cases := []struct {
		tid                int
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
	}{
		{1, "customer created Successfully", "7", []byte(res), http.StatusOK},
		{2, "customer already exists", "1", []byte(res), http.StatusBadRequest},
	}

	for x, v := range cases {
		req := httptest.NewRequest(http.MethodPost, "http://customer", bytes.NewReader(v.body))
		w := httptest.NewRecorder()

		h := handler{mockService{id: v.tid}}
		h.Create(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("Test[%v] Failed\n desc: %v\nExpected: %v \tGot: %v", x, v.desc, v.expectedStatusCode, w.Code)
		}
	}
}

func TestGetByID(t *testing.T) {
	cases := []struct {
		tid                int
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
	}{
		{1, "customer exists", "1", nil, http.StatusOK},
		{2, "customer does not exist", "100", nil, http.StatusNotFound},
	}

	for i, v := range cases {
		req := httptest.NewRequest(http.MethodGet, "http://customer", nil)
		r := mux.SetURLVars(req, map[string]string{"id": v.id})
		w := httptest.NewRecorder()

		h := handler{mockService{id: v.tid}}
		h.GetByID(w, r)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("Test[%v] Failed desc: %v\nExpected: %v \tGot %v", i, v.desc, v.expectedStatusCode, w.Code)
		}
	}
}

func TestUpdateByID(t *testing.T) {
	resp := `{"id":7,"name":"Divya S","phoneNo":"9909111143","address":"HSR Bangalore"}`
	testcases := []struct {
		tid                int
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
	}{
		{1, "customer updated successfully", "7", []byte(resp), http.StatusOK},
		{2, "customer doesn't exist", "70", []byte(resp), http.StatusInternalServerError},
	}

	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodPut, "http://customer", bytes.NewReader(v.body))
		r := mux.SetURLVars(req, map[string]string{"id": v.id})
		w := httptest.NewRecorder()

		h := handler{mockService{id: v.tid}}
		h.UpdateByID(w, r)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("Test[%v] Failed \ndesc: %v \nExpected: %v \tGot: %v", v.desc, i, v.expectedStatusCode, w.Code)
		}
	}
}

func TestDeleteByID(t *testing.T) {
	cases := []struct {
		tid                int
		desc               string
		id                 string
		body               []byte
		expectedStatusCode int
	}{
		{1, "customer deleted successfully", "7", nil, http.StatusOK},
		{2, "customer record doesn't exist", "16", nil, http.StatusInternalServerError},
	}

	for i, v := range cases {
		req := httptest.NewRequest(http.MethodDelete, "http://customer", nil)
		r := mux.SetURLVars(req, map[string]string{"id": v.id})
		w := httptest.NewRecorder()

		h := handler{mockService{id: v.tid}}
		h.DeleteByID(w, r)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != v.expectedStatusCode {
			t.Errorf("Test[%v] Failed \ndesc: %v\nExpected: %v \tGot: %v", i, v.desc, v.expectedStatusCode, w.Code)
		}

	}
}

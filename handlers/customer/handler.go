package customer

import (
	"customers-api/models"
	"customers-api/stores"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"net/http"
)

type handler struct {
	store stores.Customer
}

func New(store stores.Customer) handler {
	return handler{store: store}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var c models.Customer
	if err := ctx.Bind(&c); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if c.ID == "0" || len(c.Name) == 0 || len(c.PhoneNo) < 10 || len(c.Address) == 0 {
		return nil, errors.InvalidParam{Param: []string{c.ID, c.Name, c.PhoneNo, c.Address}}
	}
	result, err := h.store.Create(ctx, c)
	if err != nil {
		return nil, errors.DB{Err: errors.Error("internal server error")}
	}

	resp := models.Response{
		Customer:   result,
		Message:    "Created Successfully",
		StatusCode: http.StatusCreated,
	}
	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	resp, err := h.store.GetByID(ctx, id)

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "Customer",
			ID:     id,
		}
	}

	result := models.Response{}
	result = models.Response{
		Customer:   resp,
		Message:    "Retrieved Customer Successfully",
		StatusCode: http.StatusOK,
	}

	return result, nil

}

func (h handler) UpdateByID(ctx *gofr.Context) (interface{}, error) {
	var c models.Customer
	id := ctx.PathParam("id")

	if err := ctx.Bind(&c); err != nil {
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if len(id) == 0 {
		return nil, errors.InvalidParam{Param: []string{id}}
	}

	err := h.store.UpdateByID(ctx, id, c)
	if err != nil {
		return nil, errors.DB{Err: errors.Error("internal server error")}
	}

	result := models.Response{
		Customer:   c,
		Message:    "Updated Successfully",
		StatusCode: http.StatusOK,
	}
	return result, nil
}

func (h handler) DeleteByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if len(id) == 0 {
		return nil, errors.InvalidParam{Param: []string{id}}
	}

	_, err := h.store.DeleteByID(ctx, id)
	if err != nil {
		return nil, errors.DB{Err: errors.Error("internal server error")}
	}

	resp := models.Response{
		Customer:   models.Customer{ID: id},
		Message:    "Deleted Successfully",
		StatusCode: http.StatusOK,
	}
	return resp, nil
}

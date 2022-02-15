package models

type Response struct {
	Customer   interface{} `json:"customer"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statuscode"`
}

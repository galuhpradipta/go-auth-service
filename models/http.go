package models

type (
	HttpResponse struct {
		Error string      `json:"error"`
		Data  interface{} `json:"data"`
	}
)

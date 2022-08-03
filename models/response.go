package models

type (
	PageResponse struct {
		Message string
		Result  interface{}
		Status  int
	}
)

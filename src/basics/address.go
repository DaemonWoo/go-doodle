package basics

import (
	geo "url-checker/src/basics/coordinates"
)

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city" validate:"required"`
	State      string `json:"stated"`
	PostalCode string `json:"postal_code"`
	geo.Coordinates
}

// value reciever
func (a Address) isLA() bool {
	return a.City == "LA"
}

type Rectangle struct {
	Width, Height float64
}

// pointer reciever
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

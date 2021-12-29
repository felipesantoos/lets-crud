package models

import (
	"letscrud/src/api/endpoints/dto/response"
)

type Customer struct {
	Id        int
	CPF       string
	Name      string
	BirthDate string
}

func (c Customer) ConvertToDTO() response.CustomerResponse {
	return response.CustomerResponse{
		Id:        c.Id,
		CPF:       c.CPF,
		Name:      c.Name,
		BirthDate: c.BirthDate,
	}
}

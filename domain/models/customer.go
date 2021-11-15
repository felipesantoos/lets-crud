package models

import "letscrud/endpoints/response"

type Customer struct {
	Id        int64
	CPF       string
	Name      string
	BirthDate string
}

func (c Customer) ConvertToDTO() response.CustomerResponse {
	return response.CustomerResponse{
		Id:        c.Id,
		CPF:       c.CPF,
		Name:      c.Name,
		Birthdate: c.BirthDate,
	}
}

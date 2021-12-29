package tests

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/api/endpoints/dto/response"
	"letscrud/src/domain/models"
)

func GetValidCustomerRequestForService() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithInvalidCPF() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-45",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithEmptyName() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithNameWithOnlyOneWord() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "Lyara",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithNameWithSpecialCharacters() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "Lyar Caparica @nofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithNameWithNumber() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "Lyar Caparica 0nofre",
		BirthDate: "2001-01-01",
	}
}

func GetReturnedCustomerModelList() []models.Customer {
	return []models.Customer{
		{
			Id:        1,
			CPF:       "09862956046",
			Name:      "Lyara Caparica Onofre",
			BirthDate: "2001-01-01",
		},
		{
			Id:        2,
			CPF:       "78045161000",
			Name:      "Christopher Graça Sacramento",
			BirthDate: "2002-02-02",
		},
	}
}

func GetExpectedCustomerResponseList() []response.CustomerResponse {
	return []response.CustomerResponse{
		{
			Id:        1,
			CPF:       "098.629.560-46",
			Name:      "Lyara Caparica Onofre",
			BirthDate: "2001-01-01",
		},
		{
			Id:        2,
			CPF:       "780.451.610-00",
			Name:      "Christopher Graça Sacramento",
			BirthDate: "2002-02-02",
		},
	}
}

func GetCustomerRequestWithBirthDateBadFormattedForService() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "01/01/2001",
	}
}

func GetReturnedCustomerModel() *models.Customer {
	return &models.Customer{
		Id:        1,
		CPF:       "09862956046",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetExpectedCustomerResponse() *response.CustomerResponse {
	return &response.CustomerResponse{
		Id:        1,
		CPF:       "098.629.560-46",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestUpdatedForService() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "169.644.000-95",
		Name:      "Ivanilson Fernandes Baptista",
		BirthDate: "2003-03-03",
	}
}

func GetCustomerRequestUpdatedWithEmptyName() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestUpdatedWithNameWithOnlyOneWord() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "Lyara",
		BirthDate: "2001-01-01",
	}
}

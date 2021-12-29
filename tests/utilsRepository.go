package tests

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/domain/models"
)

func GetValidCustomerRequestForRepository() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "09862956046",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithCPFTooLongForRepository() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098629560460",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithBirthDateBadFormatted() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "09862956046",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "01/01/2001",
	}
}

func GetExpectedCustomerModelList() []models.Customer {
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

func GetExpectedCustomerModelWithId1() *models.Customer {
	return &models.Customer{
		Id:        1,
		CPF:       "09862956046",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetExpectedCustomerModelWithId2() *models.Customer {
	return &models.Customer{
		Id:        2,
		CPF:       "78045161000",
		Name:      "Christopher Graça Sacramento",
		BirthDate: "2002-02-02",
	}
}

func GetCustomerRequestUpdated() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "16964400095",
		Name:      "Ivanilson Fernandes Baptista",
		BirthDate: "2003-03-03",
	}
}

func GetCustomerRequestUpdatedWithCPFTooLong() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "169.644.000-95",
		Name:      "Ivanilson Fernandes Baptista",
		BirthDate: "2003-03-03",
	}
}

func GetCustomerRequestUpdatedWithBirthDateBadFormatted() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "16964400095",
		Name:      "Ivanilson Fernandes Baptista",
		BirthDate: "03/03/2003",
	}
}

func GetCustomerRequestUpdatedWithCPFAlreadyRegistered() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "78045161000",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestUpdatedWithIdenticalData() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "09862956046",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

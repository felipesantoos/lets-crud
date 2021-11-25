package tests

import (
	"letscrud/data/db"
	"letscrud/domain/models"
	"letscrud/endpoints/dto/request"
)

const CLEAR_CUSTOMER_TABLE = "DELETE FROM customer WHERE id != 0"
const RESET_AUTO_INCREMET = "ALTER TABLE customer AUTO_INCREMENT = 0"
const INSERT_CUSTOMER_ID_1 = `
	INSERT INTO customer (id, cpf, name, birthDate) VALUES 
	(1, "09862956046", "Lyara Caparica Onofre", "2001-01-01")
`
const INSERT_CUSTOMER_ID_2 = `
	INSERT INTO customer (id, cpf, name, birthDate) VALUES
	(2, "78045161000", "Christopher Graça Sacramento", "2002-02-02")
`

func SetUp(queries []string) {
	conn, err := db.GetConnection()
	if err != nil {
		panic(err)
	}

	for _, query := range queries {
		_, err = conn.Exec(query)
		if err != nil {
			panic(err)
		}
	}

	err = conn.Close()
	if err != nil {
		panic(err)
	}
}

func GetValidCustomerRequestForRepository() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "09862956046",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCustomerRequestWithCPFTooLong() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
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

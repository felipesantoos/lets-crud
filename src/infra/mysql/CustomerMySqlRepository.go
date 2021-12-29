package mysql

import (
	"fmt"
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/domain/errs"
	"letscrud/src/domain/models"
	"log"
)

type CustomerMySqlRepository struct{}

func NewCustomerMySqlRepository() *CustomerMySqlRepository {
	return &CustomerMySqlRepository{}
}

func (cr CustomerMySqlRepository) CreateNewCustomer(customerRequest request.CustomerRequest) (int, *errs.ApiError) {
	log.Println("R [CreateNewCustomer]")

	conn, err := getConnection()
	if err != nil {
		log.Println("R [CreateNewCustomer]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return 0, apiError
	}
	defer closeConnection(conn)

	sql := "INSERT INTO customer (cpf, name, birthDate) VALUES (?, ?, ?)"

	res, err := conn.Exec(sql, customerRequest.CPF, customerRequest.Name, customerRequest.BirthDate)
	if err != nil {
		log.Println("R [CreateNewCustomer]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return 0, apiError
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.Println("R [CreateNewCustomer]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return 0, apiError
	}

	id := int(lastInsertId)

	return id, nil
}

func (cr CustomerMySqlRepository) ReadAllCustomers() ([]models.Customer, *errs.ApiError) {
	log.Println("R [ReadAllCustomers]")

	conn, err := getConnection()
	if err != nil {
		log.Println("R [ReadAllCustomers]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}
	defer closeConnection(conn)

	sql := "SELECT * FROM customer"

	res, err := conn.Query(sql)
	if err != nil {
		log.Println("R [ReadAllCustomers]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}

	customerList := []models.Customer{}

	for res.Next() {
		var customer models.Customer

		err = res.Scan(&customer.Id, &customer.CPF, &customer.Name, &customer.BirthDate)
		if err != nil {
			log.Println("R [ReadAllCustomers]: " + err.Error())
			apiError := errs.GetCustomerError(err)

			return nil, apiError
		}

		customerList = append(customerList, customer)
	}

	if len(customerList) == 0 {
		apiError := errs.NewBadRequestError("Nenhum registro foi retornado!")

		return nil, apiError
	}

	return customerList, nil
}

func (cr CustomerMySqlRepository) ReadCustomerById(id int) (*models.Customer, *errs.ApiError) {
	log.Println("R [ReadCustomerById]")

	conn, err := getConnection()
	if err != nil {
		log.Println("R [ReadCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}
	defer closeConnection(conn)

	sql := "SELECT * FROM customer WHERE id = ?"
	res, err := conn.Query(sql, id)
	if err != nil {
		log.Println("R [ReadCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}

	customer := new(models.Customer)

	for res.Next() {
		err := res.Scan(&customer.Id, &customer.CPF, &customer.Name, &customer.BirthDate)
		if err != nil {
			log.Println("R [ReadCustomerById]: " + err.Error())
			apiError := errs.GetCustomerError(err)

			return nil, apiError
		}
	}

	if customer.Id == 0 {
		apiError := errs.NewBadRequestError("O cliente informado não foi encontrado!")

		return nil, apiError
	}

	return customer, nil
}

func (cr CustomerMySqlRepository) UpdateCustomerById(id int, customerRequest request.CustomerRequest) (bool, *errs.ApiError) {
	log.Println("R [UpdateCustomerById]")

	conn, err := getConnection()
	if err != nil {
		log.Println("R [UpdateCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return false, apiError
	}
	defer closeConnection(conn)

	sql := "UPDATE customer SET CPF = ?, name = ?, birthDate = ? WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		log.Println("R [UpdateCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return false, apiError
	}

	res, err := stmt.Exec(customerRequest.CPF, customerRequest.Name, customerRequest.BirthDate, id)
	if err != nil {
		log.Println("R [UpdateCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		log.Println(stmt)
		return false, apiError
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("R [UpdateCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return false, apiError
	}

	if rowsAffected == 0 {
		errorNotFound := "O cliente informado não foi encontrado"
		errorIdenticalData := "os dados informados são idênticos aos já cadastrados!"
		errorMessage := fmt.Sprintf("%s ou %s", errorNotFound, errorIdenticalData)

		apiError := errs.NewBadRequestError(errorMessage)

		return false, apiError
	}

	return true, nil
}

func (cr CustomerMySqlRepository) DeleteCustomerById(id int) (bool, *errs.ApiError) {
	log.Println("R [DeleteCustomerById]")

	conn, err := getConnection()
	if err != nil {
		log.Println("R [DeleteCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return false, apiError
	}
	defer closeConnection(conn)

	sql := "DELETE FROM customer WHERE id = ?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		log.Println("R [DeleteCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return false, apiError
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Println("R [DeleteCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return false, apiError
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("R [DeleteCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return false, apiError
	}

	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

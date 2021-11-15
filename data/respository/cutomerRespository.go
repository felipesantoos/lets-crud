package repository

import (
	"letscrud/data/db"
	"letscrud/domain/errs"
	"letscrud/domain/models"
	"letscrud/endpoints/request"
	"log"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (cr CustomerRepository) CreateNewCustomer(customer request.CustomerRequest) (int64, *errs.ApiError) {
	log.Println("Repository: CreateNewCustomer")

	conn, err := db.GetConnection()
	if err != nil {
		log.Println("Repository (CreateNewCustomer): " + err.Error())
		apiError := errs.GetCustomerError(err)

		return 0, apiError
	}
	defer conn.Close()

	sql := "INSERT INTO customer (cpf, name, birthDate) VALUES (?, ?, ?)"

	res, err := conn.Exec(sql, customer.CPF, customer.Name, customer.BirthDate)
	if err != nil {
		log.Println("Repository (CreateNewCustomer): " + err.Error())
		apiError := errs.GetCustomerError(err)

		return 0, apiError
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.Println("Repository (CreateNewCustomer): " + err.Error())
		apiError := errs.GetCustomerError(err)

		return 0, apiError
	}

	return lastInsertId, nil
}

func (cr CustomerRepository) ReadAllCustomers() ([]models.Customer, *errs.ApiError) {
	log.Println("Repository: ReadAllCustomers")

	conn, err := db.GetConnection()
	if err != nil {
		log.Println("Repository (ReadAllCustomers): " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}

	sql := "SELECT * FROM customer"

	res, err := conn.Query(sql)
	if err != nil {
		log.Println("Repository (ReadAllCustomers): " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}
	defer res.Close()

	customerList := []models.Customer{}

	for res.Next() {
		var customer models.Customer

		err = res.Scan(&customer.Id, &customer.CPF, &customer.Name, &customer.BirthDate)
		if err != nil {
			log.Println("Repository (ReadAllCustomers): " + err.Error())
			apiError := errs.GetCustomerError(err)

			return nil, apiError
		}

		customerList = append(customerList, customer)
	}

	return customerList, nil
}

func (cr CustomerRepository) ReadCustomerById() {
	log.Println("Repository: ReadCustomerById")
}

func (cr CustomerRepository) UpdateCustomerById() {
	log.Println("Repository: UpdateCustomerById")
}

func (cr CustomerRepository) DeleteCustomerById() {
	log.Println("Repository: DeleteCustomerById")
}

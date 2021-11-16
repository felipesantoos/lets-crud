package repository

import (
	"letscrud/data/db"
	"letscrud/domain/errs"
	"letscrud/domain/models"
	"letscrud/endpoints/dto/request"
	"log"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (cr CustomerRepository) CreateNewCustomer(customerRequest request.CustomerRequest) (int64, *errs.ApiError) {
	log.Println("R [CreateNewCustomer]")

	conn, err := db.GetConnection()
	if err != nil {
		log.Println("R [CreateNewCustomer]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return 0, apiError
	}
	defer conn.Close()

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

	return lastInsertId, nil
}

func (cr CustomerRepository) ReadAllCustomers() ([]models.Customer, *errs.ApiError) {
	log.Println("R [ReadAllCustomers]")

	conn, err := db.GetConnection()
	if err != nil {
		log.Println("R [ReadAllCustomers]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}

	sql := "SELECT * FROM customer"

	res, err := conn.Query(sql)
	if err != nil {
		log.Println("R [ReadAllCustomers]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}
	defer res.Close()

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

	return customerList, nil
}

func (cr CustomerRepository) ReadCustomerById(id int64) (*models.Customer, *errs.ApiError) {
	log.Println("R [ReadCustomerById]")

	con, err := db.GetConnection()
	if err != nil {
		log.Println("R [ReadCustomerById]: " + err.Error())
		apiError := errs.GetCustomerError(err)

		return nil, apiError
	}

	sql := "SELECT * FROM customer WHERE id = ?"
	res, err := con.Query(sql, id)
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

	return customer, nil
}

func (cr CustomerRepository) UpdateCustomerById(id int64, customerRequest request.CustomerRequest) {
	log.Println("R [UpdateCustomerById]")
	log.Println(id, customerRequest)
}

func (cr CustomerRepository) DeleteCustomerById() {
	log.Println("R [DeleteCustomerById]")
}

package repository

import (
	"letscrud/data/db"
	"letscrud/domain/dtos"
	"letscrud/domain/models"
	"log"
)

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (cr CustomerRepository) CreateNewCustomer(customer dtos.CustomerDTO) int64 {
	log.Println("Repository: CreateNewCustomer")

	conn, err := db.GetConnection()
	if err != nil {
		log.Println("Repository (CreateNewCustomer): " + err.Error())
	}
	defer conn.Close()

	sql := "INSERT INTO customer (cpf, name, birthDate) VALUES (?, ?, ?)"

	res, err := conn.Exec(sql, customer.CPF, customer.Name, customer.BirthDate)
	if err != nil {
		log.Println("Repository (CreateNewCustomer): " + err.Error())
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.Println("Repository (CreateNewCustomer): " + err.Error())
	}

	return lastInsertId
}

func (cr CustomerRepository) ReadAllCustomers() {
	log.Println("Repository: ReadAllCustomers")

	conn, err := db.GetConnection()
	if err != nil {
		log.Println("Repository (ReadAllCustomers): " + err.Error())
	}

	sql := "SELECT * FROM customer"

	res, err := conn.Query(sql)
	if err != nil {
		log.Println("Repository (ReadAllCustomers): " + err.Error())
	}
	defer res.Close()

	for res.Next() {
		var customer models.Customer

		err = res.Scan(&customer.Id, &customer.CPF, &customer.Name, &customer.BirthDate)
		if err != nil {
			log.Println("Repository (ReadAllCustomers): " + err.Error())
		}

		log.Println(customer)
	}
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

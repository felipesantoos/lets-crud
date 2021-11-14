package repository

import "log"

type CustomerRepository struct{}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (cr CustomerRepository) CreateNewCustomer() {
	log.Println("Repository: CreateNewCustomer")
}

func (cr CustomerRepository) ReadAllCustomers() {
	log.Println("Repository: ReadAllCustomers")
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

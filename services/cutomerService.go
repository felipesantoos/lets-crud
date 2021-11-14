package services

import (
	"letscrud/data/interfaces"
	repository "letscrud/data/respository"
	"log"
)

type CustomerService struct {
	repository interfaces.ICustomerRepository
}

func NewCustomerService() *CustomerService {
	repository := repository.NewCustomerRepository()

	return &CustomerService{repository: repository}
}

func (cs CustomerService) CreateNewCustomer() {
	log.Println("Service: CreateNewCustomer")

	cs.repository.CreateNewCustomer()
}

func (cs CustomerService) ReadAllCustomers() {
	log.Println("Service: ReadAllCustomers")

	cs.repository.ReadAllCustomers()
}

func (cs CustomerService) ReadCustomerById() {
	log.Println("Service: ReadCustomerById")

	cs.repository.ReadCustomerById()
}

func (cs CustomerService) UpdateCustomerById() {
	log.Println("Service: UpdateCustomerById")

	cs.repository.UpdateCustomerById()
}

func (cs CustomerService) DeleteCustomerById() {
	log.Println("Service: DeleteCustomerById")

	cs.repository.DeleteCustomerById()
}

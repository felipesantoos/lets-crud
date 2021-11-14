package services

import (
	"letscrud/data/interfaces"
	repository "letscrud/data/respository"
	"letscrud/domain/dtos"
	"log"
	"strings"
)

type CustomerService struct {
	repository interfaces.ICustomerRepository
}

func NewCustomerService() *CustomerService {
	repository := repository.NewCustomerRepository()

	return &CustomerService{repository: repository}
}

func (cs CustomerService) CreateNewCustomer(customer dtos.CustomerDTO) int64 {
	log.Println("Service: CreateNewCustomer")

	customer.CPF = strings.Replace(customer.CPF, ".", "", -1)
	customer.CPF = strings.Replace(customer.CPF, "-", "", -1)

	lastInsertId, _ := cs.repository.CreateNewCustomer(customer)

	return lastInsertId
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

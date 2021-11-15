package services

import (
	"letscrud/data/interfaces"
	repository "letscrud/data/respository"
	"letscrud/domain/errs"
	"letscrud/endpoints/request"
	"letscrud/endpoints/response"
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

func (cs CustomerService) CreateNewCustomer(customer request.CustomerRequest) (int64, *errs.ApiError) {
	log.Println("Service: CreateNewCustomer")

	customer.CPF = strings.Replace(customer.CPF, ".", "", -1)
	customer.CPF = strings.Replace(customer.CPF, "-", "", -1)

	lastInsertId, apiErr := cs.repository.CreateNewCustomer(customer)

	return lastInsertId, apiErr
}

func (cs CustomerService) ReadAllCustomers() ([]response.CustomerResponse, *errs.ApiError) {
	log.Println("Service: ReadAllCustomers")

	customerList, apiErr := cs.repository.ReadAllCustomers()
	if apiErr != nil {
		return nil, apiErr
	}

	var responseList []response.CustomerResponse
	for _, customer := range customerList {
		responseList = append(responseList, customer.ConvertToDTO())
	}

	return responseList, apiErr
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

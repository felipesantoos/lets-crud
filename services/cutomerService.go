package services

import (
	"letscrud/data/interfaces"
	repository "letscrud/data/respository"
	"letscrud/domain/errs"
	"letscrud/endpoints/dto/request"
	"letscrud/endpoints/dto/response"
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
	log.Println("S [CreateNewCustomer]")

	customer.CPF = strings.Replace(customer.CPF, ".", "", -1)
	customer.CPF = strings.Replace(customer.CPF, "-", "", -1)

	lastInsertId, apiErr := cs.repository.CreateNewCustomer(customer)

	return lastInsertId, apiErr
}

func (cs CustomerService) ReadAllCustomers() ([]response.CustomerResponse, *errs.ApiError) {
	log.Println("S [ReadAllCustomers]")

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

func (cs CustomerService) ReadCustomerById(id int64) (*response.CustomerResponse, *errs.ApiError) {
	log.Println("S [ReadCustomerById]")

	customer, apiError := cs.repository.ReadCustomerById(id)
	if apiError != nil {
		return nil, apiError
	}

	response := customer.ConvertToDTO()

	return &response, nil
}

func (cs CustomerService) UpdateCustomerById(id int64, customerRequest request.CustomerRequest) {
	log.Println("Service: UpdateCustomerById")

	cs.repository.UpdateCustomerById(id, customerRequest)
}

func (cs CustomerService) DeleteCustomerById() {
	log.Println("Service: DeleteCustomerById")

	cs.repository.DeleteCustomerById()
}

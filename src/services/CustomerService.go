package services

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/api/endpoints/dto/response"
	"letscrud/src/domain/errs"
	interfaces "letscrud/src/services/repositories"
	"letscrud/src/services/utils"
	"log"
	"strings"
)

type CustomerService struct {
	repository interfaces.ICustomerRepository
}

func NewCustomerService(repository interfaces.ICustomerRepository) *CustomerService {
	return &CustomerService{repository: repository}
}

func (cs CustomerService) CreateNewCustomer(customerRequest request.CustomerRequest) (int64, *errs.ApiError) {
	log.Println("S [CreateNewCustomer]")

	customerRequest.CPF = utils.FormatCPFRemovingPunctuation(customerRequest.CPF)

	if !utils.IsValidCPF(customerRequest.CPF) {
		return 0, errs.NewBadRequestError("O CPF informado é inválido!")
	}

	if !strings.Contains(strings.TrimSpace(customerRequest.Name), " ") {
		return 0, errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")
	}

	if !utils.IsValidName(customerRequest.Name) {
		return 0, errs.NewBadRequestError("O nome só pode conter letras e espaços!")
	}

	lastInsertedId, apiErr := cs.repository.CreateNewCustomer(customerRequest)

	return lastInsertedId, apiErr
}

func (cs CustomerService) ReadAllCustomers() ([]response.CustomerResponse, *errs.ApiError) {
	log.Println("S [ReadAllCustomers]")

	customerList, apiErr := cs.repository.ReadAllCustomers()
	if apiErr != nil {
		return nil, apiErr
	}

	var responseList []response.CustomerResponse
	for _, customer := range customerList {
		customer.CPF = utils.FormatCPFAddingPunctuation(customer.CPF)

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

	customer.CPF = utils.FormatCPFAddingPunctuation(customer.CPF)

	response := customer.ConvertToDTO()

	return &response, nil
}

func (cs CustomerService) UpdateCustomerById(id int64, customerRequest request.CustomerRequest) (bool, *errs.ApiError) {
	log.Println("Service: UpdateCustomerById")

	customerRequest.CPF = utils.FormatCPFRemovingPunctuation(customerRequest.CPF)

	if !utils.IsValidCPF(customerRequest.CPF) {
		return false, errs.NewBadRequestError("O CPF informado é inválido!")
	}

	if !strings.Contains(strings.TrimSpace(customerRequest.Name), " ") {
		return false, errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")
	}

	isUpdated, apiError := cs.repository.UpdateCustomerById(id, customerRequest)
	if apiError != nil {
		return false, apiError
	}

	return isUpdated, nil
}

func (cs CustomerService) DeleteCustomerById(id int64) (bool, *errs.ApiError) {
	log.Println("Service: DeleteCustomerById")

	isDeleted, apiError := cs.repository.DeleteCustomerById(id)
	if apiError != nil {
		return false, apiError
	}

	return isDeleted, nil
}

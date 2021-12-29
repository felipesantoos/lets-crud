package interfaces

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/domain/errs"
	"letscrud/src/domain/models"
)

type ICustomerRepository interface {
	CreateNewCustomer(customerRequest request.CustomerRequest) (int, *errs.ApiError)
	ReadAllCustomers() ([]models.Customer, *errs.ApiError)
	ReadCustomerById(id int) (*models.Customer, *errs.ApiError)
	UpdateCustomerById(id int, customerRequest request.CustomerRequest) (bool, *errs.ApiError)
	DeleteCustomerById(id int) (bool, *errs.ApiError)
}

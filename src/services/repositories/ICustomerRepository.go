package interfaces

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/domain/errs"
	"letscrud/src/domain/models"
)

type ICustomerRepository interface {
	CreateNewCustomer(customerRequest request.CustomerRequest) (int64, *errs.ApiError)
	ReadAllCustomers() ([]models.Customer, *errs.ApiError)
	ReadCustomerById(id int64) (*models.Customer, *errs.ApiError)
	UpdateCustomerById(id int64, customerRequest request.CustomerRequest) (bool, *errs.ApiError)
	DeleteCustomerById(id int64) (bool, *errs.ApiError)
}

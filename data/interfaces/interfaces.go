package interfaces

import (
	"letscrud/domain/errs"
	"letscrud/domain/models"
	"letscrud/endpoints/request"
)

type ICustomerRepository interface {
	CreateNewCustomer(request.CustomerRequest) (int64, *errs.ApiError)
	ReadAllCustomers() ([]models.Customer, *errs.ApiError)
	ReadCustomerById()
	UpdateCustomerById()
	DeleteCustomerById()
}

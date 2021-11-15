package interfaces

import (
	"letscrud/domain/errs"
	"letscrud/endpoints/request"
	"letscrud/endpoints/response"
)

type ICustomerService interface {
	CreateNewCustomer(request.CustomerRequest) (int64, *errs.ApiError)
	ReadAllCustomers() ([]response.CustomerResponse, *errs.ApiError)
	ReadCustomerById()
	UpdateCustomerById()
	DeleteCustomerById()
}

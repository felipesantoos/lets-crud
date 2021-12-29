package interfaces

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/api/endpoints/dto/response"
	"letscrud/src/domain/errs"
)

type ICustomerService interface {
	CreateNewCustomer(customerResquest request.CustomerRequest) (int, *errs.ApiError)
	ReadAllCustomers() ([]response.CustomerResponse, *errs.ApiError)
	ReadCustomerById(id int) (*response.CustomerResponse, *errs.ApiError)
	UpdateCustomerById(id int, customerRequest request.CustomerRequest) (bool, *errs.ApiError)
	DeleteCustomerById(id int) (bool, *errs.ApiError)
}

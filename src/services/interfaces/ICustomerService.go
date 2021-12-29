package interfaces

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/api/endpoints/dto/response"
	"letscrud/src/domain/errs"
)

type ICustomerService interface {
	CreateNewCustomer(customerResquest request.CustomerRequest) (int64, *errs.ApiError)
	ReadAllCustomers() ([]response.CustomerResponse, *errs.ApiError)
	ReadCustomerById(id int64) (*response.CustomerResponse, *errs.ApiError)
	UpdateCustomerById(id int64, customerRequest request.CustomerRequest) (bool, *errs.ApiError)
	DeleteCustomerById(id int64) (bool, *errs.ApiError)
}

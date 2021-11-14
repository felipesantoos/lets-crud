package interfaces

import (
	"letscrud/domain/dtos"
)

type ICustomerService interface {
	CreateNewCustomer(dtos.CustomerDTO) int64
	ReadAllCustomers()
	ReadCustomerById()
	UpdateCustomerById()
	DeleteCustomerById()
}

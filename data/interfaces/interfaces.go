package interfaces

import (
	"letscrud/domain/dtos"
)

type ICustomerRepository interface {
	CreateNewCustomer(dtos.CustomerDTO) int64
	ReadAllCustomers()
	ReadCustomerById()
	UpdateCustomerById()
	DeleteCustomerById()
}

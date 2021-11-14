package interfaces

type ICustomerService interface {
	CreateNewCustomer()
	ReadAllCustomers()
	ReadCustomerById()
	UpdateCustomerById()
	DeleteCustomerById()
}

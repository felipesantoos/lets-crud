package interfaces

type ICustomerRepository interface {
	CreateNewCustomer()
	ReadAllCustomers()
	ReadCustomerById()
	UpdateCustomerById()
	DeleteCustomerById()
}

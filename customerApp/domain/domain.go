package domain

type Customer struct {
	Id int
	Name string
	Email string
}

type CustomerStore interface {
	CreateCustomer(Customer) error
	DeleteCustomer(int) error
	GetAllCustomer() ([]Customer ,error)
	GetCustomerById(int) (Customer, error)
	UpdateCustomer(int, Customer) error
}
package entities

type Customer struct {
	Id   int
	Name string
}

type CustomerRepository interface {
	Store(customer Customer) error
	FindById(id int) Customer
}

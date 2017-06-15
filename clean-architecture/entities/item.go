package entities

type Item struct {
	Id        int
	Name      string
	Value     float64
	Available bool
}

type ItemRepository interface {
	Store(item Item) error
	FindById(id int) Item
}

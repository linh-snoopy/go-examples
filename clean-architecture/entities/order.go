package entities

import (
	"errors"
)

var MAXIMUM = 250.00

type Order struct {
	Id       int
	Customer Customer
	Items    []Item
}

type OrderRepository interface {
	Store(order Order) error
	FindById(id int) Order
}

func (order *Order) value() float64 {
	sum := 0.0
	for i := range order.Items {
		sum += order.Items[i].Value
	}
	return sum
}

func (order *Order) Add(item Item) error {
	if !item.Available {
		return errors.New("Can not add unavailable items to order.")
	}
	if order.value()+item.Value > MAXIMUM {
		return errors.New(`An order may not exceed a total value of 250.00`)
	}
	order.Items = append(order.Items, item)
	return nil
}

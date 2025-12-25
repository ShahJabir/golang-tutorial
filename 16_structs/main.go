package main

import (
	"fmt"
	"time"
)

type Customer struct {
	ID        int
	Name      string
	Phone     string
	CreatedAt time.Time
}

type Order struct {
	ID     int
	Amount float64
	Status string
	Customer
	CreatedAt time.Time
}

func NewCustomer(id int, name, phone string) *Customer {
	return &Customer{
		ID:        id,
		Name:      name,
		Phone:     phone,
		CreatedAt: time.Now(),
	}
}

func NewOrder(id int, amount float64, status string, customer Customer) *Order {
	return &Order{
		ID:        id,
		Amount:    amount,
		Status:    status,
		Customer:  customer,
		CreatedAt: time.Now(),
	}
}

func (o *Order) changeOrderStatus(status string) {
	o.Status = status
}

func (o Order) getOrderTime() time.Time {
	return o.CreatedAt
}

func main() {
	customer := NewCustomer(1, "John Doe", "123-456-7890")
	fmt.Println("Customer:", customer)
	order := NewOrder(1, 99.99, "Pending", *customer)
	fmt.Println("Order:", order)
	order.changeOrderStatus("Shipped")
	fmt.Println("Updated Order:", order)
	fmt.Println("Order Created At:", order.getOrderTime())

	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println("Language Struct:", language)

}

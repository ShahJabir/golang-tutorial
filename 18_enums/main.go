package main

import "fmt"

type OrderStatus string

const (
	Pending    OrderStatus = "Pending"
	Processing OrderStatus = "Processing"
	Shipped    OrderStatus = "Shipped"
	Delivered  OrderStatus = "Delivered"
)

func ChangeOrderStatus(status OrderStatus) {
	fmt.Println("Order status:", status)
}

func main() {
	ChangeOrderStatus(Pending)
	ChangeOrderStatus(Processing)
	ChangeOrderStatus(Shipped)
	ChangeOrderStatus(Delivered)
}

package main

import "fmt"

type paymenter interface {
	pay(amount float64)
}

type payment struct {
	gateway paymenter
}

type stripe struct{}

type fakePaymentGateway struct{}

func (p payment) makePayment(amount float64) {
	p.gateway.pay(amount)
}

func (s stripe) pay(amount float64) {
	fmt.Printf("Paid %.2f using Stripe\n", amount)
}

func (f fakePaymentGateway) pay(amount float64) {
	fmt.Printf("Pretending to pay %.2f using FakePaymentGateway\n", amount)
}

func main() {
	stripeGateway := stripe{}
	fakeGateway := fakePaymentGateway{}
	payment1 := payment{gateway: stripeGateway}
	payment2 := payment{gateway: fakeGateway}
	payment1.makePayment(100.0)
	payment2.makePayment(50.0)
	fmt.Println("Payments processed successfully")
}

package main

import "fmt"

type PaymentService struct {
	method PaymentMethod
}

type PaymentMethod interface {
	Pay(amount float64)
}

type Bkash struct {
	apiKey string
}

type Nagad struct {
	apiKey string
}

func (bk Bkash) Pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Bkash", amount)
}

func (ng Nagad) Pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Nagad", amount)
}

func (ps PaymentService) checkout() {
	ps.method.Pay(100.00)
}

func main() {
	// bkash := Bkash{apiKey: "1234"}
	nagad := Nagad{apiKey: "1234"}

	ps := PaymentService{
		method: nagad,
	}
	ps.checkout()
}

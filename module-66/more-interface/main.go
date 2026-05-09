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

func (bk *Bkash) Pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Bkash \n", amount)
}

func (ng *Nagad) Pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Nagad \n", amount)
}

func (ps PaymentService) checkout(amount float64) {
	ps.method.Pay(amount)
}

func NewPaymentServie(method PaymentMethod) *PaymentService {
	return &PaymentService{method}
}

func NewBkash(apiKey string) *Bkash {
	return &Bkash{apiKey: apiKey}
}

func NewNagad(apiKey string) *Nagad {
	return &Nagad{apiKey: apiKey}
}

func main() {
	bkash := NewBkash("1234")
	payment1 := NewPaymentServie(bkash)
	payment1.checkout(100.00)

	nagad := NewNagad("1234")
	payment2 := NewPaymentServie(nagad)
	payment2.checkout(500.00)

	mockPM := MockPaymentMethod{}
	mockPayment := NewPaymentServie(&mockPM)
	mockPayment.checkout(1000.00)
}

type MockPaymentMethod struct{}

func (m *MockPaymentMethod) Pay(amount float64) {
	fmt.Printf("Testing Payment Method %.2f tk....!", amount)
}

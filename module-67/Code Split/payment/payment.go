package payment

import "fmt"

type paymentService struct {
	method paymentMethod
}

type paymentMethod interface {
	Pay(amount float64)
}

type bkash struct {
	apiKey string
}

type nagad struct {
	apiKey string
}

func (bk *bkash) Pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Bkash \n", amount)
}

func (ng *nagad) Pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Nagad \n", amount)
}

func (ps paymentService) Checkout(amount float64) {
	ps.method.Pay(amount)
}

func NewPaymentServie(method paymentMethod) *paymentService {
	return &paymentService{method}
}

func NewBkash(apiKey string) *bkash {
	return &bkash{apiKey: apiKey}
}

func NewNagad(apiKey string) *nagad {
	return &nagad{apiKey: apiKey}
}
package testing

import "fmt"

type MockPaymentMethod struct{}

func (m *MockPaymentMethod) Pay(amount float64) {
	fmt.Printf("Testing Payment Method %.2f tk....!", amount)
}

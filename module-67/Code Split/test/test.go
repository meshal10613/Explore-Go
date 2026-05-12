package testing

import "fmt"

func init(){
	fmt.Println("Test Package init...")
}

type MockPaymentMethod struct{}

func (m *MockPaymentMethod) Pay(amount float64) {
	fmt.Printf("Testing Payment Method %.2f tk....!", amount)
}

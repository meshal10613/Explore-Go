package main

import (
	"CodeSplit/payment"
	testing "CodeSplit/test" //? alias

	"github.com/fatih/color"
)

//? package name smaller -> private e.g. checkout
//? package name bigger -> public e.g. Checkout

func main() {
	bkash := payment.NewBkash("1234")
	payment1 := payment.NewPaymentServie(bkash)
	payment1.Checkout(100.00)

	nagad := payment.NewNagad("1234")
	payment2 := payment.NewPaymentServie(nagad)
	payment2.Checkout(500.00)

	mockPM := testing.MockPaymentMethod{}
	mockPayment := payment.NewPaymentServie(&mockPM)
	mockPayment.Checkout(1000.00)

	color.New(color.BgGreen, color.FgBlue).Println("\nSuccess")
}

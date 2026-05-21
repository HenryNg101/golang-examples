package main

import "fmt"

// Strategy type by function instead of interface
type PaymentStrategy func(amount float64)

// Implementations of strats
func CreditCardPay(amount float64) {
	fmt.Println("Paid with credit card:", amount)
}
func PayPalPay(amount float64) {
	fmt.Println("Paid with PayPal:", amount)
}

// Use the strategy in client struct
type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p PaymentProcessor) Process(amount float64) {
	p.strategy(amount)
}

func main() {
	// Dynamically use it in code
	processor := PaymentProcessor{strategy: CreditCardPay}
	processor.Process(100)
	processor.strategy = PayPalPay
	processor.Process(200)
}

package main

import (
	"fmt"
)

// Payer represents any type that can make a payment.
type Payer interface {
	Pay(amount int) error
}

// Wallet represents a simple wallet with available cash.
type Wallet struct {
	Cash int
}

// Pay deducts money from the wallet if there are enough funds.
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("insufficient funds")
	}

	w.Cash -= amount
	return nil
}

// Buy processes a purchase using any payment method
// that implements the Payer interface.
func Buy(p Payer) {
	if err := p.Pay(10); err != nil {
		panic(err)
	}

	fmt.Printf("Thank you for your purchase using %T\n", p)
}

func main() {
	myWallet := &Wallet{
		Cash: 100,
	}

	Buy(myWallet)
}

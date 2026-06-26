package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("not enough cash")
	}

	w.Cash -= amount
	return nil
}

// String implements the fmt.Stringer interface.
func (w *Wallet) String() string {
	return "Wallet containing " + strconv.Itoa(w.Cash) + " units of cash"
}

func main() {
	myWallet := &Wallet{
		Cash: 100,
	}

	fmt.Printf("Raw wallet: %#v\n", myWallet)
	fmt.Printf("Payment method: %s\n", myWallet)

	// Equivalent:
	fmt.Println(myWallet)
}

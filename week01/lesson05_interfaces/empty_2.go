package main

import (
	"fmt"
	"strconv"
)

// -------------------------------------------------------

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
	return "wallet containing " + strconv.Itoa(w.Cash) + " units of cash"
}

// -------------------------------------------------------

type Payer interface {
	Pay(amount int) error
}

// -------------------------------------------------------

func Buy(in any) {
	p, ok := in.(Payer)
	if !ok {
		fmt.Printf("%T is not a payment method\n\n", in)
		return
	}

	if err := p.Pay(10); err != nil {
		fmt.Printf("payment failed with %T: %v\n\n", p, err)
		return
	}

	fmt.Printf("Thank you for your purchase using %T\n\n", p)
}

// -------------------------------------------------------

func main() {
	myWallet := &Wallet{
		Cash: 100,
	}

	Buy(myWallet)
	Buy([]int{1, 2, 3})
	Buy(3.14)
}

package main

import "fmt"

// --------------

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

// --------------

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("not enough money on balance")
	}

	c.Balance -= amount
	return nil
}

// --------------

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("not enough money on account")
	}

	a.Money -= amount
	return nil
}

// --------------

type Payer interface {
	Pay(amount int) error
}

// --------------

func Buy(p Payer) {
	switch paymentMethod := p.(type) {
	case *Wallet:
		fmt.Println("Payment by cash")
	case *Card:
		fmt.Println("Insert your card,", paymentMethod.Cardholder)
	default:
		fmt.Println("New payment method")
	}

	if err := p.Pay(10); err != nil {
		fmt.Printf("payment failed with %T: %v\n\n", p, err)
		return
	}

	fmt.Printf("Thank you for your purchase using %T\n\n", p)
}

// --------------

func main() {
	myWallet := &Wallet{
		Cash: 100,
	}

	Buy(myWallet)

	var myMoney Payer

	myMoney = &Card{
		Balance:    100,
		Cardholder: "Andrew Mart",
	}

	Buy(myMoney)

	myMoney = &ApplePay{
		Money: 9,
	}

	Buy(myMoney)
}

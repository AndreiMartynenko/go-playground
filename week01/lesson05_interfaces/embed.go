package main

import (
	"fmt"
)

type Phone struct {
	Money   int
	AppleID string
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("not enough money on the account")
	}

	p.Money -= amount
	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("please enter a phone number")
	}

	return nil
}

// -------------------------------------------------------

type Payer interface {
	Pay(amount int) error
}

type Ringer interface {
	Ring(number string) error
}

// NFCPhone combines the Payer and Ringer interfaces.
type NFCPhone interface {
	Payer
	Ringer
}

// -------------------------------------------------------

// PayForMetroWithPhone accepts any type
// that implements both Payer and Ringer.
func PayForMetroWithPhone(phone NFCPhone) {
	if err := phone.Pay(1); err != nil {
		fmt.Printf("payment failed: %v\n", err)
		return
	}

	fmt.Printf("The gate has been opened using %T\n", phone)
}

// -------------------------------------------------------

func main() {
	myPhone := &Phone{
		Money: 9,
	}

	PayForMetroWithPhone(myPhone)
}

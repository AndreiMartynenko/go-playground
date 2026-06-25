package main

//type BankAccount struct {
//	ID      int
//	Owner   string
//	Balance float64
//}
//
//func (b BankAccount) String() string {
//	return fmt.Sprintf("%s %s (%d)\n: %d", b.ID, b.Owner, b.Balance)
//}
//
//func (b *BankAccount) Deposit(amount float64) {
//	b.Balance += amount
//	fmt.Println("Your balance is: ", b.Balance)
//}
//
//func (b *BankAccount) Withdraw(amount float64) error {
//	if b.Balance < amount {
//		return fmt.Errorf("insufficient funds")
//
//	}
//	b.Balance -= amount
//	return nil
//
//}
//
//// Transfer
//func (b *BankAccount) Transfer(to *BankAccount, amount float64) error {
//	if to == nil {
//		return fmt.Errorf("destination account is nil")
//	}
//
//	if amount <= 0 {
//		return fmt.Errorf("amount must be greater than zero")
//	}
//
//	if b.Balance < amount {
//		return fmt.Errorf("insufficient funds")
//	}
//
//	b.Balance -= amount
//	to.Balance += amount
//
//	return nil
//}
//
//func main() {
//	bankAccount := BankAccount{
//		ID:      1,
//		Owner:   "John Doe",
//		Balance: 1000,
//	}
//
//	fmt.Println(bankAccount.String())
//
//}

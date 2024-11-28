package bank

import "fmt"

type BankAccount interface {
	Deposit(float64) string
	Withdraw(float64) string
	CheckBalance() float64
}

type SavingsAccount struct {
	balance float64
}
type CheckingAccount struct {
	balance float64
}
type LoanAccount struct {
	balance float64
}

func (s *SavingsAccount) Deposit(amount float64) string {
	fmt.Printf("Your saving account replenished by amount %v$ \n", amount)
	s.balance += amount
	return "Saving account deposit"
}
func (s *SavingsAccount) Withdraw(amount float64) string {
	if amount > s.balance {
		fmt.Printf("Insufficient funds in your saving accounts\n")
		return "Operation cancelled"
	}
	fmt.Printf("From your saving account withdrawn %v$ \n", amount)
	s.balance -= amount
	return "Success operation"
}
func (s *SavingsAccount) CheckBalance() float64 {
	fmt.Printf("Balance of your saving account is %v$ \n", s.balance)
	return s.balance
}

func (s *CheckingAccount) Deposit(amount float64) string {
	fmt.Printf("Your checking account replenished by amount %v$ \n", amount)
	s.balance += amount
	return "Checking account deposit"
}
func (s *CheckingAccount) Withdraw(amount float64) string {
	if amount > s.balance {
		fmt.Printf("Insufficient funds in your checking account\n")
		return "Operation cancelled"
	}
	fmt.Printf("From your checking account withdrawn %v$ \n", amount)
	s.balance -= amount
	return "Success operation"
}
func (s *CheckingAccount) CheckBalance() float64 {
	fmt.Printf("Balance of your checking account is %v$ \n", s.balance)
	return s.balance
}

func (l *LoanAccount) Deposit(amount float64) string {
	fmt.Printf("Your loan account replenished by amount %v$ \n", amount)
	l.balance += amount
	return "Loan account deposit"
}
func (l *LoanAccount) Withdraw(amount float64) string {
	if amount > l.balance {
		fmt.Printf("Insufficient funds in your loan account\n")
		return "Operation cancelled"
	}
	fmt.Printf("From your loan account withdrawn %v$ \n", amount)
	l.balance -= amount
	return "Success operation"
}
func (s *LoanAccount) CheckBalance() float64 {
	fmt.Printf("Balance of your loan account is %v$ \n", s.balance)
	return s.balance
}

func ReplenishAnyAccount(a BankAccount, amount float64) {
	a.Deposit(amount)
}
func WithdrawFromAccount(a BankAccount, amount float64) {
	a.Withdraw(amount)
}
func ShowAllBalances(a []BankAccount) {
	for _, account := range a {
		account.CheckBalance()
	}
}

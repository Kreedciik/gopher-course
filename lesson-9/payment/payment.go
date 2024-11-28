package payment

import "fmt"

type PaymentProcessor interface {
	Pay(float64) string
	Refund(float64) string
	TransactionHistory() []string
}

type CreditCardProcessor struct {
	CardNumber   string
	Transactions []string
}

type PaypalProcessor struct {
	Email        string
	Transactions []string
}

type CryptoProcessor struct {
	WalletAddress string
	Transactions  []string
}

// CreditCardProcessor methods
func (c *CreditCardProcessor) Pay(amount float64) string {
	fmt.Printf("Paying by credit card with card number %v", c.CardNumber)
	transaction := fmt.Sprintf("Registered transaction with amount of %0.2f", amount)
	c.Transactions = append(c.Transactions, transaction)
	return transaction
}

func (c *CreditCardProcessor) Refund(amount float64) string {
	fmt.Printf("Refunded amount of %v to this card %v", amount, c.CardNumber)
	transaction := fmt.Sprintf("Registered refunding with amount of %0.2f", amount)
	c.Transactions = append(c.Transactions, transaction)
	return transaction
}

func (c *CreditCardProcessor) TransactionHistory() []string {
	fmt.Println("Transaction history of credit card")
	return c.Transactions
}

// PaypalProcessor methods
func (p *PaypalProcessor) Pay(amount float64) string {
	fmt.Printf("Paying by Paypal account with Email %v", p.Email)
	transaction := fmt.Sprintf("Registered transaction with amount of %0.2f", amount)
	p.Transactions = append(p.Transactions, transaction)
	return transaction
}

func (p *PaypalProcessor) Refund(amount float64) string {
	fmt.Printf("Refunded amount of %v to this account with Email %v", amount, p.Email)
	transaction := fmt.Sprintf("Registered refunding with amount of %0.2f", amount)
	p.Transactions = append(p.Transactions, transaction)
	return transaction
}

func (p *PaypalProcessor) TransactionHistory() []string {
	fmt.Println("Transaction history of Paypal account")
	return p.Transactions
}

// CryptoProcessor methods
func (c *CryptoProcessor) Pay(amount float64) string {
	fmt.Printf("Paying by crypto wallet with crypto address %v", c.WalletAddress)
	transaction := fmt.Sprintf("Registered transaction with amount of %0.2f", amount)
	c.Transactions = append(c.Transactions, transaction)
	return transaction
}

func (c *CryptoProcessor) Refund(amount float64) string {
	fmt.Printf("Refunded amount of %v to this wallet %v", amount, c.WalletAddress)
	transaction := fmt.Sprintf("Registered refunding with amount of %0.2f", amount)
	c.Transactions = append(c.Transactions, transaction)
	return transaction
}

func (c *CryptoProcessor) TransactionHistory() []string {
	fmt.Println("Transaction history of crypto wallet")
	return c.Transactions
}

func ShowTransactions(p []PaymentProcessor) {
	for _, v := range p {
		v.TransactionHistory()
	}
}

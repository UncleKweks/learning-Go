package main

import (
	"fmt"
)

type BankAccount struct {
	accountNumber string
	accountName   string
	balance       float64
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Deposit amount must be greater than zero")
	}
	b.balance += amount
	fmt.Printf("Deposited %.2f to account %s. New balance: %.2f\n", amount, b.accountNumber, b.balance)
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Withdrawal amount must be greater than zero")
	}
	if amount > b.balance {
		return fmt.Errorf("Insufficient funds")
	}
	b.balance -= amount
	fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, b.accountNumber, b.balance)
	return nil
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func (b *BankAccount) GetAccountDetails() string {
	return fmt.Sprintf("Account Number: %s, Account Name: %s, Balance: %.2f", b.accountNumber, b.accountName, b.balance)
}

type SavingsAccount struct {
	BankAccount
	interestRate float64
}

func (s *SavingsAccount) ApplyInterest() {
	interest := s.balance * s.interestRate / 100
	fmt.Printf("Applied interest of %.2f to account %s. New balance: %.2f\n", interest, s.accountNumber, s.balance)
	err := s.Deposit(interest)
	if err != nil {
		fmt.Println("Error applying interest:", err)
	}
	s.balance += interest
}

type OverdraftAccount struct {
	BankAccount
	overdraftLimit float64
}

func (o *OverdraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Withdrawal amount must be greater than zero")
	}
	if amount > o.balance+o.overdraftLimit {
		return fmt.Errorf("Insufficient funds")
	}
	o.balance -= amount
	fmt.Printf("Withdrew %.2f from account %s. New balance: %.2f\n", amount, o.accountNumber, o.balance)
	return nil
}

func main() {

	fmt.Println("Bank Account Management System")

	savAccount := SavingsAccount{
		BankAccount: BankAccount{
			accountNumber: "SA123456",
			accountName:   "John Doe",
			balance:       1000.00,
		},
		interestRate: 2.5,
	}

	fmt.Println(savAccount.GetAccountDetails())

	err := savAccount.Deposit(500)
	if err != nil {
		fmt.Println("Error:", err)
	}
	savAccount.ApplyInterest()

	err = savAccount.Withdraw(200)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(savAccount.GetAccountDetails())

	overdraftAccount := OverdraftAccount{
		BankAccount: BankAccount{
			accountNumber: "OD123456",
			accountName:   "Jane Smith",
			balance:       500.00,
		},
		overdraftLimit: 300.00,
	}

	fmt.Println(overdraftAccount.GetAccountDetails())

	err = overdraftAccount.Withdraw(700)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(overdraftAccount.GetAccountDetails())
	
	finalBalance := overdraftAccount.GetBalance()
	fmt.Printf("Final balance for account %s: %.2f\n", overdraftAccount.accountNumber, finalBalance)

}

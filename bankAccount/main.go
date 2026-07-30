package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (b *BankAccount) Balance() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.balance
}

func (b *BankAccount) Withdraw(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.balance < amount {
		fmt.Println("Cannot Withdraw that amount", amount)
		return
	}

	b.balance -= amount
	fmt.Println("Withdraw", amount)
}

func (b *BankAccount) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.balance < amount {
		fmt.Println("Cannot Deposit that amount", amount)
	}

	b.balance += amount
	fmt.Println("Deposit", amount)
}

func main() {

	var wg sync.WaitGroup
	var account = &BankAccount{balance: 200}

	//for i := 0; i <= 10; i++ {
	//	wg.Add(1)
	//
	//	go func(amount int) {
	//		defer wg.Done()
	//		time.Sleep(time.Duration(amount) * time.Millisecond)
	//		account.Deposit(amount)
	//	}(i + 1)

	for i := 0; i <= 6; i++ {
		wg.Add(1)

		go func(amount int) {
			defer wg.Done()
			time.Sleep(time.Duration(amount) * time.Millisecond)
			account.Withdraw(amount * 10)
		}(i + 1)
	}
	//wg.Wait()
	fmt.Println(account.Balance())
}

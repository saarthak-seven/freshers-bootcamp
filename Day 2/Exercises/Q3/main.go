package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type BankAccount struct {
	balance int64
	mu      sync.Mutex
}

func (acc *BankAccount) Deposit(amount int) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if amount <= 0 {
		return fmt.Errorf("Enter a valid amount")
	}

	acc.balance += int64(amount)
	return nil
}

func (acc *BankAccount) Withdraw(amount int) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if amount > int(acc.balance) {

		return fmt.Errorf("Not enough Balance")
	}
	acc.balance -= int64(amount)
	return nil
}

func (acc *BankAccount) Balance() int64 {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	return acc.balance
}

func NewBankAccount(amount int64) *BankAccount {
	return &BankAccount{
		balance: amount,
	}
}

func main() {
	account := BankAccount{
		balance: 500,
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			err := account.Deposit(rand.Intn(1000) + i)
			if err != nil {
				fmt.Println("Deposit failed:", err)
			}
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			err := account.Withdraw(rand.Intn(1000))
			if err != nil {
				fmt.Println("Withdraw failed:", err)
			}
		}()
	}
	wg.Wait()

	fmt.Println(account.balance)
}

package main

import (
	"fmt"
	"sync"
)

// Counter with mutex protection
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value safely reads the counter value
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	fmt.Println("=== Mutex Example ===\n")

	fmt.Println("1. WITHOUT Mutex:")
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()

	fmt.Printf("Counter = %d\n", counter)
	// return
	// Example 2: Safe access with mutex
	fmt.Println("2. WITH Mutex:")
	muCounter := &Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			muCounter.Increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Counter = %d\n", muCounter.Value())
	// Example 3: Multiple operations protected by mutex
	fmt.Println("3. Mutex protecting multiple operations:")
	account := &BankAccount{balance: 0}

	// Concurrent deposits
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(amount float64) {
			defer wg.Done()
			account.Deposit(amount)
			fmt.Printf("   Deposited %.2f, new balance: %.2f\n", amount, account.Balance())
		}(10.0)
	}
	wg.Wait()

	fmt.Printf("\n   Final balance: %.2f\n", account.Balance())
}

// BankAccount demonstrates mutex usage for more complex operations
type BankAccount struct {
	mu      sync.Mutex
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

func (b *BankAccount) Balance() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

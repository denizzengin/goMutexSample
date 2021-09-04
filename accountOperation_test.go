package accountOperation

import (
	"testing"
	"sync"
)

func TestAccountOperation(t *testing.T) {
	var wg sync.WaitGroup
	amount := 1000
	wg.Add(2) // think that count of goroutines
	go Withdraw(100, amount, &wg)
	go Deposit(100, amount, &wg)
	wg.Wait() // Wait until count of goroutines is equal to 0.

	if amount != 1000 {
		t.Log("Not passed because of wrong mutex uses")
	}
}
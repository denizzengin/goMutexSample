package accountOperation

import (
	"testing"
	"sync"
	"github.com/shopspring/decimal"
)

func TestAccountOperation(t *testing.T) {
	var wg sync.WaitGroup
		
	b, err := decimal.NewFromString("1000.001") 
	if err != nil {
		t.Log("Not passed")
	}
	a := Account{ Balance : b }
	wg.Add(2) // think that count of goroutines	
	go Withdraw(decimal.RequireFromString("10.001"), &a, &wg)
	go Deposit(decimal.RequireFromString("10.001"), &a, &wg)
	wg.Wait() // Wait until count of goroutines is equal to 0.

	if !a.Balance.Equal(b) {		
		t.Errorf("Not passed, expected %s actual %s", a.Balance.StringFixed(3), b.StringFixed(3))
	}
}
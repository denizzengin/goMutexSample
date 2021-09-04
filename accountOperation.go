package accountOperation

import (
	"fmt"
	"sync"
	"github.com/shopspring/decimal"
)

var (
	mutex sync.Mutex
)

type Account struct {
	Balance decimal.Decimal 
}

func Withdraw(amountOfWithdraw decimal.Decimal, a *Account, wg *sync.WaitGroup) {
	mutex.Lock()	
	a.Balance = a.Balance.Sub(amountOfWithdraw)
	mutex.Unlock()
	wg.Done(); // Decrease 1 value count of goroutines
	fmt.Printf("New balance is %s \n", a.Balance.StringFixed(3))
}

func Deposit(amountOfDeposit decimal.Decimal, a *Account, wg *sync.WaitGroup) {
	mutex.Lock()
	a.Balance = a.Balance.Add(amountOfDeposit)
	mutex.Unlock()
	wg.Done() // Decrease 1 value count of goroutines
	fmt.Printf("New balance is %s \n", a.Balance.StringFixed(3))
}


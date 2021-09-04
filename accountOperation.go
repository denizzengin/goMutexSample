package accountOperation

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
)
func Withdraw(amountOfWithdraw int, amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	amount -= amountOfWithdraw
	mutex.Unlock()
	wg.Done(); // Decrease 1 value count of goroutines
	fmt.Printf("New balance is %d \n", amount)
}

func Deposit(amountOfDeposit int, amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	amount += amountOfDeposit
	mutex.Unlock()
	wg.Done() // Decrease 1 value count of goroutines
	fmt.Printf("New balance is %d \n", amount)
}


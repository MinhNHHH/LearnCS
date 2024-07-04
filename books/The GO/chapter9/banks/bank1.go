// Package bank provides a concurrency-safe bank with one account.
package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrawals = make(chan WithdrawRequest)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	fmt.Println("Start")
	for {
		select {
		case amount := <-deposits:
			balance += amount
			fmt.Println("Start1")
		case w := <-withdrawals:
			if w.Amount > balance {
				w.Success <- false
				continue
			}
			fmt.Println("Start2")
			balance -= w.Amount
			w.Success <- true
		case balances <- balance:
		}
	}
}

type WithdrawRequest struct {
	Success chan bool
	Amount  int
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdrawals <- WithdrawRequest{Success: ch, Amount: amount}
	return <-ch
}

func Init() {
	go teller() // start the monitor goroutine
}

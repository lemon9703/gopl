// Package bank provides a concurrency-safe bank with one account.
package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

type withdrawal struct {
	amount  int
	succeed chan bool
}

var withdraws = make(chan withdrawal)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	succeed := make(chan bool)
	withdraws <- withdrawal{amount, succeed}
	return <-succeed
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdrawal := <-withdraws:
			if withdrawal.amount <= balance {
				balance -= withdrawal.amount
				withdrawal.succeed <- true
			} else {
				withdrawal.succeed <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

package pointeranderror

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

//interface defined in the fmt package and lets you define how your type is printed
//when used with %s format string in prints

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// var allows declaring global values
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

//lowercase varaible, function etc makes the variable private
//all arguments, instances method is called on are pass by value
//struct pointers are automatically derefrenced

//by convention, all method reciever type should be same for concistency

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"gopl-solutions/Jave/solutions/ch9/ex9-1"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// Jave withdraw success
	go func() {
		bank.Withdraw(150)
		done <- struct{}{}
	}()
	<-done
	if got, want := bank.Balance(), 150; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	// Jave withdraw fail
	go func() {
		bank.Withdraw(200)
		done <- struct{}{}
	}()
	<-done
	if got, want := bank.Balance(), 150; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

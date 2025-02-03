package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(20)

	fmt.Printf("address of wallet in test is %v\n", &wallet.balance)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

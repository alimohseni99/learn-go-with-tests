package pointers

import (
	"testing"
)

func TestWallet(t *testing.T){

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		t.Helper()
		got := wallet.Balance()
		if got != want {
				t.Errorf("got %d want %d", got, want)
		}
	}
	assertError := func(t testing.TB, err error)  {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	assertBalance(t,wallet,  Bitcoin(10))
	
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance:  Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(25)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})

}


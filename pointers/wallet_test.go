package pointers

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit 10 bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw all money", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertErrInsufficientFunds(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertErrInsufficientFunds(t testing.TB, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	var err InsufficientFundsErr
	if !errors.As(got, &err) {
		t.Fatal("want InsufficientFunds but didn't get one")
	}

	t.Log(err)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

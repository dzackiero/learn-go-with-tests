package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()

		if got == nil && want != nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit bitcoin to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw bitcoin from wallet", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(5)
		want := Bitcoin(5)

		assertBalance(t, wallet, want)
		assertError(t, err, nil)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(100)

		// Check if error is happened
		assertError(t, err, ErrInsufficientFunds)
		// check if the balance not changing
		assertBalance(t, wallet, startingBalance)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}

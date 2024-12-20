package pointeranderror

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(10)
		//errcheck tells that the line of code above may return an error, and we have
		//not checked for it
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		//we'd expect withdraw function to return error which is an interface which
		//can be nillable.
		assertBalance(t, wallet, startingBalance)
		//err is not nil when test is run succesfully, because we are trying to
		//genrate the error
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one!")
		//stops the test if called so that further condition isin't run
	}
	//err.Error() return string value of message
	if err != want {
		t.Errorf("got %q, want %q", err, want)
		//if we get some other type of error then this is invoked
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

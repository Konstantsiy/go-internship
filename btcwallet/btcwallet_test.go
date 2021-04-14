package btcwallet

import (
	"sync"
	"testing"
)

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("No error expected but got: %v", err)
	}
}

func assertError(t *testing.T, givenError, expectedError error) {
	t.Helper()
	if givenError != expectedError {
		t.Errorf("Eepected %v error but got another one: %v", expectedError, givenError)
	}
}

func assertBalanceEquality(t *testing.T, result, expected Bitcoin) {
	t.Helper()
	if result != expected {
		t.Errorf("Incorrect result. Expect %.3f, got %.3f", expected, result)
	}
}

func TestBtcWallet_Race(t *testing.T) {
	const (
		startBalance Bitcoin = 20
		expected     Bitcoin = 70
	)

	wallet := NewBtcWallet(startBalance)
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			err := wallet.Deposit(10)
			assertNoError(t, err)
		}()
		go func() {
			defer wg.Done()
			err := wallet.Withdraw(5)
			assertNoError(t, err)
		}()
	}

	wg.Wait()
	result := wallet.GetBalance()
	assertBalanceEquality(t, result, expected)
}

func TestBtcWallet_Deposit(t *testing.T) {
	t.Run("NoDepositErrors", func(t *testing.T) {
		wallet := NewBtcWallet(Bitcoin(23.00))

		err := wallet.Deposit(Bitcoin(23.456))
		assertNoError(t, err)

		err = wallet.Deposit(Bitcoin(0))
		assertNoError(t, err)

		expected := Bitcoin(46.456)
		result := wallet.GetBalance().RoundTo(3)
		assertBalanceEquality(t, expected, result)
	})

	t.Run("NegativeDepositAmount", func(t *testing.T) {
		wallet := NewBtcWallet(Bitcoin(10.00))

		err := wallet.Deposit(Bitcoin(-0.123))
		assertError(t, err, ErrDepositAmount)

		err = wallet.Deposit(Bitcoin(0.234))
		assertNoError(t, err)

		expected := Bitcoin(10.234)
		result := wallet.GetBalance().RoundTo(3)
		assertBalanceEquality(t, expected, result)
	})
}

func TestBtcWallet_Withdraw(t *testing.T) {
	t.Run("NoWithdrawErrors", func(t *testing.T) {
		wallet := NewBtcWallet(Bitcoin(1000))

		err := wallet.Withdraw(Bitcoin(999.445))
		assertNoError(t, err)

		expected := Bitcoin(0.555)
		result := wallet.GetBalance().RoundTo(3)
		assertBalanceEquality(t, expected, result)
	})

	t.Run("NegativeWithdrawAmount", func(t *testing.T) {
		wallet := NewBtcWallet(Bitcoin(1000))

		err := wallet.Withdraw(Bitcoin(-12.34))
		assertError(t, err, ErrWithdrawAmount)

		err = wallet.Withdraw(0.123)
		assertNoError(t, err)

		expected := Bitcoin(999.877)
		result := wallet.GetBalance().RoundTo(3)
		assertBalanceEquality(t, expected, result)
	})

	t.Run("TooLargeWithdrawAmount", func(t *testing.T) {
		wallet := NewBtcWallet(Bitcoin(1000))

		err := wallet.Withdraw(Bitcoin(1000.001))
		assertError(t, err, ErrWithdrawAmount)

		expected := Bitcoin(1000.000)
		result := wallet.GetBalance().RoundTo(3)
		assertBalanceEquality(t, result, expected)
	})
}

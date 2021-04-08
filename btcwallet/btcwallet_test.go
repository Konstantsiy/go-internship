package btcwallet

import (
	"errors"
	"testing"
)

func TestBtcWallet_Deposit(t *testing.T) {
	const (
		startBalance        Bitcoin = 20
		precision           int     = 3
		depositErrorMessage string  = "the Bitcoins amount for the deposit must be positive"
	)

	testTable := []struct {
		Source      Bitcoin
		Expected    Bitcoin
		ExpectedErr error
	}{
		{
			Source:      23.456,
			Expected:    43.456,
			ExpectedErr: nil,
		},
		{
			Source:      -0.234,
			Expected:    startBalance,
			ExpectedErr: errors.New(depositErrorMessage),
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		if err := wallet.Deposit(tc.Source); err != nil {
			if err.Error() != tc.ExpectedErr.Error() {
				t.Error(err.Error())
			}
		} else {
			result := wallet.GetBalance().RoundTo(precision)
			if result != tc.Expected {
				t.Errorf("Incorrect result. Expect %.3f, got %.3f", tc.Expected, result)
			}
		}
	}
}

func TestBtcWallet_Withdraw(t *testing.T) {
	const (
		startBalance         Bitcoin = 1000
		precision            int     = 3
		withdrawErrorMessage string  = "the Bitcoins amount for the withdraw must be limited by the balance"
	)

	testTable := []struct {
		Source      Bitcoin
		Expected    Bitcoin
		ExpectedErr error
	}{
		{
			Source:      999.445,
			Expected:    0.555,
			ExpectedErr: nil,
		},
		{
			Source:      -12.34,
			Expected:    startBalance,
			ExpectedErr: errors.New(withdrawErrorMessage),
		},
		{
			Source:      1000.34,
			Expected:    startBalance,
			ExpectedErr: errors.New(withdrawErrorMessage),
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		if err := wallet.Withdraw(tc.Source); err != nil {
			if err.Error() != tc.ExpectedErr.Error() {
				t.Error(err.Error())
			}
		} else {
			result := wallet.GetBalance().RoundTo(precision)
			if result != tc.Expected {
				t.Errorf("Incorrect result. Expect %.3f, got %.3f", tc.Expected, result)
			}
		}
	}
}

package btcwallet

import (
	"errors"
	"sync"
	"testing"
)

func TestBtcWallet_Race(t *testing.T) {
	var startBalance Bitcoin = 20
	var expected Bitcoin = 50
	wallet := NewBtcWallet(startBalance)
	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(w *BtcWallet) {
			defer wg.Done()
			for j := 0; j < 2; j++ {
				var err error
				if err = w.Deposit(10); err != nil {
					t.Errorf(err.Error())
				}
				if err = w.Withdraw(5); err != nil {
					t.Errorf(err.Error())
				}
			}
		}(wallet)
	}
	wg.Wait()
	result := wallet.GetBalance()

	if expected != result {
		t.Errorf("Incorrect result. Expect %.2f, got %.2f", expected, result)
	}
}

func TestBtcWallet_Deposit(t *testing.T) {
	const (
		startBalance        Bitcoin = 20
		precision           int     = 3
		depositErrorMessage string  = "the Bitcoins amount for the deposit must be positive"
	)

	testTable := []struct {
		SumToDeposit  Bitcoin
		Expected      Bitcoin
		ExpectedError error
	}{
		{
			SumToDeposit:  23.456,
			Expected:      43.456,
			ExpectedError: nil,
		},
		{
			SumToDeposit:  -0.234,
			Expected:      startBalance,
			ExpectedError: errors.New(depositErrorMessage),
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		if err := wallet.Deposit(tc.SumToDeposit); err != nil {
			if err.Error() != tc.ExpectedError.Error() {
				errorAssertionHelper(t, err, tc.ExpectedError)
			}
		}
		result := wallet.GetBalance().RoundTo(precision)

		if result != tc.Expected {
			t.Errorf("incorrect result. Expect %.3f, got %.3f", tc.Expected, result)
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
		SumToWithdraw Bitcoin
		Expected      Bitcoin
		ExpectedError error
	}{
		{
			SumToWithdraw: 999.445,
			Expected:      0.555,
			ExpectedError: nil,
		},
		{
			SumToWithdraw: -12.34,
			Expected:      startBalance,
			ExpectedError: errors.New(withdrawErrorMessage),
		},
		{
			SumToWithdraw: 1000.34,
			Expected:      startBalance,
			ExpectedError: errors.New(withdrawErrorMessage),
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		if err := wallet.Withdraw(tc.SumToWithdraw); err != nil {
			if err.Error() != tc.ExpectedError.Error() {
				errorAssertionHelper(t, err, tc.ExpectedError)
			}
		}
		result := wallet.GetBalance().RoundTo(precision)

		if result != tc.Expected {
			t.Errorf("incorrect result. Expect %.3f, got %.3f", tc.Expected, result)
		}
	}
}

func errorAssertionHelper(t *testing.T, givenError, expectedError error) {
	t.Helper()
	if givenError.Error() != expectedError.Error() {
		t.Error("expected error but didn't get one")
	}
}

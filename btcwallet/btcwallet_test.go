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
		startBalance Bitcoin = 0
		expected     Bitcoin = 50
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
	const (
		startBalance Bitcoin = 20
		precision    int     = 3
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
			ExpectedError: ErrDepositAmount,
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		if err := wallet.Deposit(tc.SumToDeposit); err != nil {
			assertError(t, err, tc.ExpectedError)
		}

		result := wallet.GetBalance().RoundTo(precision)
		assertBalanceEquality(t, result, tc.Expected)
	}
}

func TestBtcWallet_Withdraw(t *testing.T) {
	const (
		startBalance Bitcoin = 1000
		precision    int     = 3
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
			ExpectedError: ErrWithdrawAmount,
		},
		{
			SumToWithdraw: 1000.34,
			Expected:      startBalance,
			ExpectedError: ErrWithdrawAmount,
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		if err := wallet.Withdraw(tc.SumToWithdraw); err != nil {
			assertError(t, err, tc.ExpectedError)
		}

		result := wallet.GetBalance().RoundTo(precision)
		assertBalanceEquality(t, result, tc.Expected)
	}
}

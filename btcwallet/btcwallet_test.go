package btcwallet

import (
	"testing"
)

func TestBtcWallet_Deposit(t *testing.T) {
	const (
		startBalance Bitcoin = 0
		precision    int     = 2
	)

	testTable := []struct {
		Source   []Bitcoin
		Expected Bitcoin
	}{
		{
			Source:   []Bitcoin{2.5, 20, 30, 45.5, 38},
			Expected: 136,
		},
		{
			Source:   []Bitcoin{12, -12.45, 10.45},
			Expected: 22.45,
		},
		{
			Source:   []Bitcoin{-12, -3.45, 0},
			Expected: 0,
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		for _, btcAmount := range tc.Source {
			_ = wallet.Deposit(btcAmount)
		}
		result := wallet.GetBalance().RoundTo(precision)

		if result != tc.Expected {
			t.Errorf("Incorrect result. Expect %.2f, got %.2f", tc.Expected, result)
		}
	}
}

func TestBtcWallet_Withdraw(t *testing.T) {
	const (
		startBalance Bitcoin = 1000
		precision    int     = 2
	)

	testTable := []struct {
		Source   []Bitcoin
		Expected Bitcoin
	}{
		{
			Source:   []Bitcoin{0.34, 450, 500, 34.24, 10.2},
			Expected: 5.22,
		},
		{
			Source:   []Bitcoin{0, -12.45, 1000.01},
			Expected: 1000,
		},
		{
			Source:   []Bitcoin{999, 0.5, 0.49},
			Expected: 0.01,
		},
	}

	for _, tc := range testTable {
		wallet := NewBtcWallet(startBalance)

		for _, btcAmount := range tc.Source {
			_ = wallet.Withdraw(btcAmount)
		}
		result := wallet.GetBalance().RoundTo(precision)

		if result != tc.Expected {
			t.Errorf("Incorrect result. Expect %.2f, got %.2f", tc.Expected, result)
		}
	}
}

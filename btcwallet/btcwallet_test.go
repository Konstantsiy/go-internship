package btcwallet

import (
	"sync"
	"testing"

	_ "golang.org/x/tools/go/expect"
)

func TestBtcWallet(t *testing.T) {
	wallet := NewBtcWallet(20)
	wg := &sync.WaitGroup{}
	var expected Bitcoin = 50
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

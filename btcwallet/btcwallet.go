package btcwallet

import (
	"sync"

	"github.com/pkg/errors"
)

// Bitcoin representation via the float64 data type.
type Bitcoin float64

// BtcWallet implements a simple bitcoins wallet for deposit, withdraw and get current balance.
type BtcWallet struct {
	sync.Mutex         // Embedded sync.Mutex that safe wallet for concurrent deposit and withdrawal operations.
	balance    Bitcoin // balance stores the current value of the bitcoin(float64) balance in the wallet.
}

// NewBtcWallet returns a new BtcWallet which stores a predefined balance value.
func NewBtcWallet(balance Bitcoin) *BtcWallet {
	return &BtcWallet{balance: balance}
}

// Deposit increments the wallet balance by the given argument.
func (wallet *BtcWallet) Deposit(data Bitcoin) error {
	wallet.Lock()
	defer wallet.Unlock()

	if data <= 0 {
		return errors.New("the data for the deposit must be positive")
	}
	wallet.balance += data

	return nil
}

// Withdraw decrements the wallet balance by the given argument.
func (wallet *BtcWallet) Withdraw(data Bitcoin) error {
	wallet.Lock()
	defer wallet.Unlock()

	if data <= 0 || data > wallet.balance {
		return errors.New("the data for the withdraw must be limited by the balance")
	}
	wallet.balance -= data

	return nil
}

// GetBalance returns a current balance.
func (wallet *BtcWallet) GetBalance() Bitcoin {
	return wallet.balance
}

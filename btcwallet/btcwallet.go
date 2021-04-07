// Package btcwallet implements Bitcoins wallet for deposit, withdraw and get wallet's balance.
package btcwallet

import (
	"errors"
	"math"
	"sync"
)

// Bitcoin representation via the float64 data type.
type Bitcoin float64

// RoundTo truncates float64 type to a particular precision and return new Bitcoin.
func (b Bitcoin) RoundTo(precision int) Bitcoin {
	output := math.Pow(10, float64(precision))
	number := float64(b) * output
	rounded := int(number + math.Copysign(0.5, number))
	return Bitcoin(float64(rounded) / output)
}

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
func (wallet *BtcWallet) Deposit(amount Bitcoin) error {
	wallet.Lock()
	defer wallet.Unlock()

	if amount <= 0 {
		return errors.New("the Bitcoins amount for the deposit must be positive")
	}
	wallet.balance += amount

	return nil
}

// Withdraw decrements the wallet balance by the given argument.
func (wallet *BtcWallet) Withdraw(amount Bitcoin) error {
	wallet.Lock()
	defer wallet.Unlock()

	if amount <= 0 || amount > wallet.balance {
		return errors.New("the Bitcoins amount for the withdraw must be limited by the balance")
	}
	wallet.balance -= amount

	return nil
}

// GetBalance returns a current balance.
func (wallet *BtcWallet) GetBalance() Bitcoin {
	return wallet.balance
}

### Package btcwallet
The package implements a simple wallet for storing Bitcoins(represented by the float64 type). The wallet supports 
the following operations:
- Deposit - increasing the wallet balance by the given Bitcoins number;
- Withdraw - increasing the wallet balance by the given Bitcoins number;
- GetBalance - getting the current balance value. 

Wallets safe for concurrent deposit and withdrawal which are synchronized with mutex.


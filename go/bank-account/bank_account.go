package account

import "sync"

type Account struct {
	balance int64
	closed  bool
	mu      sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.closed {
		return 0, false
	} else {
		return a.balance, true
	}
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed || a.balance+amount < 0 {
		return a.balance, false
	} else {
		a.balance += amount
		return a.balance, true
	}
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return a.balance, false
	} else {
		a.closed = true
		b := a.balance
		a.balance = 0
		return b, true
	}
}

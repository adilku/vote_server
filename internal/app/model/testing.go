package model

import "testing"

func TestWallet(t *testing.T) *Wallet {
	t.Helper()
	return &Wallet{
		Name: "first",
		Balance: 0,
	}
}

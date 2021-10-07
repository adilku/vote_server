package model

import "testing"

func TestWallet(t *testing.T) *Wallet {
	t.Helper()
	return &Wallet{
		ID: 100,
		Balance: 0,
	}
}

func TestWallet2(t *testing.T) *Wallet {
	t.Helper()
	return &Wallet{
		ID: 101,
		Balance: 0,
	}
}


func TestCreditBalance(t *testing.T) int {
	t.Helper()
	return 2
}

func TestDebitBalance(t *testing.T) int {
	t.Helper()
	return -2
}
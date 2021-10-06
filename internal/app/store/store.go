package store

type Store interface {
	 GetWallet() WalletRepository
}
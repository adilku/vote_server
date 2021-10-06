package store

import "github.com/adilku/vote_server/internal/app/model"

type WalletRepository interface {
	Create(wallet *model.Wallet) error
	FindByName(name string) (*model.Wallet, error)
	//ChangeById(id int, poll *model.Wallet) error
}
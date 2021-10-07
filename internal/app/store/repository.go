package store

import "github.com/adilku/vote_server/internal/app/model"

type WalletRepository interface {
	Create(wallet *model.Wallet) error
	FindById(id int) (*model.Wallet, error)
	ChangeBalance(id int, delta int) error
	Transfer(idSender int, idReceiver int, delta int) error
	//ChangeById(id int, poll *model.Wallet) error
}
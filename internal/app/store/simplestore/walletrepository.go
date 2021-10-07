package simplestore

import (
	"errors"
	"github.com/adilku/vote_server/internal/app/model"
)

type WalletRepository struct {
	store  *Store
	wallet map[int]int
}

func (r *WalletRepository) ChangeBalance(id int, delta int) error {
	old := r.wallet[id]
	r.wallet[id] = old + delta
	return nil
}

func (r *WalletRepository) Create(u *model.Wallet) error {
	u.ID = len(r.wallet)
	r.wallet[u.ID] = u.Balance
	return nil
}

func (r *WalletRepository) FindById(id int) (*model.Wallet, error) {
	u, exist := r.wallet[id]
	if !exist {
		return nil, errors.New("Not found")
	}
	return &model.Wallet{ID : id, Balance: u}, nil
}

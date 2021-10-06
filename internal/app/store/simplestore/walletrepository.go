package simplestore

import (
	"errors"
	"github.com/adilku/vote_server/internal/app/model"
)

type WalletRepository struct {
	store  *Store
	wallet map[string]*model.Wallet
}

func (r *WalletRepository) Create(u *model.Wallet) error {
	u.ID = len(r.wallet)
	r.wallet[u.Name] = u
	return nil
}

func (r *WalletRepository) FindByName(name string) (*model.Wallet, error) {
	u, exist := r.wallet[name]
	if !exist {
		return nil, errors.New("Not found")
	}
	return u, nil
}

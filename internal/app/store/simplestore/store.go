package simplestore

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store"
)

type Store struct {
	pollRepository *WalletRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) GetWallet() store.WalletRepository {
	if s.pollRepository != nil {
		return s.pollRepository
	}

	s.pollRepository = &WalletRepository{
		store:  s,
		wallet: map[string]*model.Wallet{},
	}

	return s.pollRepository
}

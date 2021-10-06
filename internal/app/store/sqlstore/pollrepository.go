package sqlstore

import (
	"github.com/adilku/vote_server/internal/app/model"
)

type WalletRepository struct {
	store *Store
}

func (r *WalletRepository) Create(u *model.Wallet) error {
	sqlStatement := `
	INSERT INTO wallets (user_name, cur_balance)
	VALUES ($1, $2)
	RETURNING id`
	err := r.store.db.QueryRow(
		sqlStatement,u.Name, u.Balance,
	).Scan(&u.ID)
	return err
}


func (r *WalletRepository) FindByName(name string) (*model.Wallet, error) {
	u := &model.Wallet{}
	if err := r.store.db.QueryRow(
		"SELECT id FROM wallets WHERE  user_name = $1",
		name,
		).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}



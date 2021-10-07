package sqlstore

import (
	"github.com/adilku/vote_server/internal/app/model"
)

type WalletRepository struct {
	store *Store
}

func (r *WalletRepository) ChangeBalance(id int, delta int) error {
	if delta > 0 {
		if err := r.store.db.QueryRow(
			"UPDATE wallets SET cur_balance = cur_balance + $2 WHERE id = $1",
			id,
			delta,
		).Err(); err != nil {
			return err
		}
	} else {
		delta = -delta
		if err := r.store.db.QueryRow(
			"UPDATE wallets SET cur_balance = cur_balance - $2 WHERE id = $1",
			id,
			delta,
		).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (r *WalletRepository) Create(u *model.Wallet) error {
	sqlStatement := `
	INSERT INTO wallets (id, cur_balance)
	VALUES ($1, $2)
	RETURNING id`
	err := r.store.db.QueryRow(
		sqlStatement,u.ID, u.Balance,
	).Scan(&u.ID)
	return err
}


func (r *WalletRepository) FindById(id int) (*model.Wallet, error) {
	u := &model.Wallet{ID : id}
	if err := r.store.db.QueryRow(
		"SELECT cur_balance FROM wallets WHERE  id = $1",
		id,
		).Scan(&u.Balance); err != nil {
		return nil, err
	}
	return u, nil
}




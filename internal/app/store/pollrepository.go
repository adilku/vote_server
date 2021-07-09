package store

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/lib/pq"
)

type PollRepository struct {
	store *Store
}

func (r *PollRepository) Create(u *model.Poll) (*model.Poll, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO polls (poll) VALUES ($1::new_type[]) RETURNING id",
		pq.Array(u.PollVars),
		).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}


func (r *PollRepository) FindById(id int) (*model.Poll, error) {
	return nil, nil
}


package sqlstore

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/lib/pq"
)

type PollRepository struct {
	store *Store
}

func (r *PollRepository) Create(u *model.Poll) error {
	return r.store.db.QueryRow(
		"INSERT INTO polls (poll) VALUES ($1::new_type[]) RETURNING id",
		pq.Array(u.PollVars),
		).Scan(&u.ID)
}


func (r *PollRepository) FindById(id int) (*model.Poll, error) {
	u := &model.Poll{}

	if err := r.store.db.QueryRow(
		"SELECT poll FROM polls WHERE  id = $1",
		id,
		).Scan(pq.Array(&u.PollVars)); err != nil {
		return nil, err
	}

	return u, nil
}



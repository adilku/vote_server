package simplestore

import (
	"errors"
	"github.com/adilku/vote_server/internal/app/model"
)

type PollRepository struct {
	store *Store
	polls map[int]*model.Poll
}

func (r *PollRepository) Create(u *model.Poll) error {
	u.ID = len(r.polls)
	r.polls[u.ID] = u
	return nil
}

func (r *PollRepository) FindById(id int) (*model.Poll, error) {
	u, err := r.polls[id]
	if !err {
		return nil, errors.New("Not found")
	}
	return u, nil
}

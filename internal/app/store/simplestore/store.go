package simplestore

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store"
)

type Store struct {
	pollRepository *PollRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) GetPoll() store.PollRepository {
	if s.pollRepository != nil {
		return s.pollRepository
	}

	s.pollRepository = &PollRepository{
		store: s,
		polls: map[int]*model.Poll{},
	}

	return s.pollRepository
}

package sqlstore

import (
	"database/sql"
	"github.com/adilku/vote_server/internal/app/store"
	_ "github.com/lib/pq"
)


type Store struct {
	db *sql.DB
	pollRepository *PollRepository
}

func New (db *sql.DB) *Store {
	return &Store{
		db : db,
	}
}


func (s *Store) GetPoll() store.PollRepository {
	if s.pollRepository != nil {
		return s.pollRepository
	}

	s.pollRepository = &PollRepository{
		store: s,
	}

	return s.pollRepository
}



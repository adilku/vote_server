package sqlstore_test

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("polls")

	s := sqlstore.New(db)
	u := model.TestPoll(t)
	assert.NoError(t, s.GetPoll().Create(u))
	assert.NotNil(t, u)
}

func TestPollRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("polls")
	s := sqlstore.New(db)
	u1 := model.TestPoll(t)
	_, err := s.GetPoll().FindById(u1.ID)
	assert.Error(t, err)

	s.GetPoll().Create(u1)
	f, err := s.GetPoll().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

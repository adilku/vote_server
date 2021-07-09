package store_test

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPollRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)

	defer teardown("polls")
	u, err := s.GetPoll().Create(&model.Poll{
		PollVars: []model.PairPoll{{Name: "1", Cnt: 0}},
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

package simplestore_test

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store/simplestore"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestPollRepository_Create(t *testing.T) {
	s := simplestore.New()
	u := model.TestPoll(t)
	assert.NoError(t, s.GetPoll().Create(u))
	assert.NotNil(t, u)
}

func TestPollRepository_FindById(t *testing.T) {
	s := simplestore.New()
	u1 := model.TestPoll(t)
	_, err := s.GetPoll().FindById(u1.ID)
	assert.Error(t, err)

	s.GetPoll().Create(u1)
	f, err := s.GetPoll().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

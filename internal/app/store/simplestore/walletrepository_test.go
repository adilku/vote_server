package simplestore_test

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store/simplestore"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestPollRepository_Create(t *testing.T) {
	s := simplestore.New()
	u := model.TestWallet(t)
	assert.NoError(t, s.GetWallet().Create(u))
	assert.NotNil(t, u)
}

func TestPollRepository_FindByName(t *testing.T) {
	s := simplestore.New()
	u1 := model.TestWallet(t)
	_, err := s.GetWallet().FindByName(u1.Name)
	assert.Error(t, err)

	s.GetWallet().Create(u1)
	f, err := s.GetWallet().FindByName(u1.Name)
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

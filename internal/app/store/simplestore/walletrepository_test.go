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
	_, err := s.GetWallet().FindById(u1.ID)
	assert.Error(t, err)

	s.GetWallet().Create(u1)
	f, err := s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f)
}

func TestPollRepository_ChangeBalance(t *testing.T)  {
	s := simplestore.New()
	u1 := model.TestWallet(t)
	oldBalance := u1.Balance
	err := s.GetWallet().ChangeBalance(u1.ID , model.TestCreditBalance(t))
	assert.NoError(t, err)
	newGoodBalance := oldBalance + model.TestCreditBalance(t)
	curWallet, err := s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, newGoodBalance, curWallet.Balance)
}

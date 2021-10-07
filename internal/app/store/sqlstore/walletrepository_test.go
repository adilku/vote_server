package sqlstore_test

import (
	"github.com/adilku/vote_server/internal/app/model"
	"github.com/adilku/vote_server/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalletRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("wallets")

	s := sqlstore.New(db)
	u := model.TestWallet(t)
	assert.NoError(t, s.GetWallet().Create(u))
	assert.NotNil(t, u)
}

func TestWalletRepository_FindById(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("wallets")
	s := sqlstore.New(db)
	u1 := model.TestWallet(t)
	_, err := s.GetWallet().FindById(u1.ID)
	assert.Error(t, err)

	s.GetWallet().Create(u1)
	f, err := s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, f)
}


func TestWalletRepository_ChangeBalance(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("wallets")
	s := sqlstore.New(db)
	u1 := model.TestWallet(t)
	s.GetWallet().Create(u1)
	oldWallet, err := s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	s.GetWallet().ChangeBalance(u1.ID, model.TestDebitBalance(t))
	newExpectBalance := oldWallet.Balance + model.TestDebitBalance(t)
	newCurBalance, err := s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, newExpectBalance, newCurBalance.Balance)

	oldWallet, err = s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	s.GetWallet().ChangeBalance(u1.ID, model.TestCreditBalance(t))
	newExpectBalance = oldWallet.Balance + model.TestCreditBalance(t)
	newCurBalance, err = s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, newExpectBalance, newCurBalance.Balance)
}



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
	err = s.GetWallet().ChangeBalance(u1.ID, model.TestCreditBalance(t))
	newExpectBalance := oldWallet.Balance + model.TestCreditBalance(t)
	newCurBalance, err := s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, newExpectBalance, newCurBalance.Balance)

	oldWallet, err = s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	s.GetWallet().ChangeBalance(u1.ID, model.TestDebitBalance(t))
	newExpectBalance = oldWallet.Balance + model.TestDebitBalance(t)
	newCurBalance, err = s.GetWallet().FindById(u1.ID)
	assert.NoError(t, err)
	assert.EqualValues(t, newExpectBalance, newCurBalance.Balance)
}

func TestWalletRepository_NegativeBalance(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("wallets")
	s := sqlstore.New(db)
	u1 := model.TestWallet(t)
	s.GetWallet().Create(u1)
	err := s.GetWallet().ChangeBalance(u1.ID, model.TestDebitBalance(t))
	assert.Error(t, err)
}

func TestWalletRepository_transfer(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("wallets")
	s := sqlstore.New(db)
	u1 := model.TestWallet(t)
	s.GetWallet().Create(u1)
	u2 := model.TestWallet2(t)
	s.GetWallet().Create(u2)
	err := s.GetWallet().Transfer(u2.ID, u1.ID, model.TestCreditBalance(t))
	assert.NoError(t, err)
	balance2, _ := s.GetWallet().FindById(u2.ID)
	assert.EqualValues(t, u2.Balance - model.TestCreditBalance(t), balance2.Balance)
}


func TestWalletRepository_badtransfer(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("wallets")
	s := sqlstore.New(db)
	u1 := model.TestWallet(t)
	s.GetWallet().Create(u1)
	u2 := model.TestWallet2(t)
	s.GetWallet().Create(u2)
	err := s.GetWallet().Transfer(u1.ID, u2.ID, model.TestCreditBalance(t))
	assert.Error(t, err)
	//balance2, _ := s.GetWallet().FindById(u2.ID)
	//assert.EqualValues(t, u2.Balance - model.TestCreditBalance(t), balance2.Balance)
}


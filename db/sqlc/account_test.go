package db

import (
	"context"
	"github.com/seanlion/golang_bank_proj/db/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	// main_test.go에서 선언한 testQueries를 여기서 사용한다.
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

// 인자로 넣을 때는 testing.T로 넣어야 함.
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	testingAccount := createRandomAccount(t)
	fetchedAccount, err := testQueries.GetAccount(context.Background(), testingAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, testingAccount.ID, fetchedAccount.ID)
	require.Equal(t, testingAccount.Owner, fetchedAccount.Owner)
	require.Equal(t, testingAccount.Balance, fetchedAccount.Balance)
	require.Equal(t, testingAccount.Currency, fetchedAccount.Currency)

	require.WithinDuration(t, testingAccount.CreatedAt, fetchedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	existingAccount := createRandomAccount(t)

	arg := UpdateAccountBalanceParams{
		ID:      existingAccount.ID,
		Balance: util.RandomMoney(),
	}

	err := testQueries.UpdateAccountBalance(context.Background(), arg)

	require.NoError(t, err)

	fetchedAccount, err := testQueries.GetAccount(context.Background(), existingAccount.ID)

	require.Equal(t, arg.Balance, fetchedAccount.Balance)
}

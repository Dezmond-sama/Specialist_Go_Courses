package db

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Dezmond-sama/Specialist_Go_Courses/GO3/bankstore/utils"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

const dbSource = "postgresql://postgres:postgres@localhost:5433/bankstoredb?sslmode=disable"

var ctx = context.Background()

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("can't connect to db", err)
	}
	defer conn.Close(ctx)

	testQueries = New(conn)
	os.Exit(m.Run())
}

func CreateRandomAccountParams() CreateAccountParams {
	return CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomAmount(),
		Currency: Currency(utils.RandomCurrency()),
	}
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}
func createRandomAccount(t *testing.T) Accounts {
	arg := CreateRandomAccountParams()
	account, err := testQueries.CreateAccount(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.Created)

	return account
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(ctx, account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.Created.Time, account2.Created.Time, time.Second)
}

func TestUpdateAccountAmount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountBalanceParams{
		ID:      account1.ID,
		Balance: utils.RandomAmount(),
	}
	account2, err := testQueries.UpdateAccountBalance(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, utils.PgNumericToFloat(arg.Balance), utils.PgNumericToFloat(account2.Balance))

	require.WithinDuration(t, account1.Created.Time, account2.Created.Time, time.Second)
}

func TestUpdateAccountOwner(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountOwnerParams{
		ID:    account1.ID,
		Owner: utils.RandomOwner(),
	}
	account2, err := testQueries.UpdateAccountOwner(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.Created.Time, account2.Created.Time, time.Second)
}
func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(ctx, account1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(ctx, account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(ctx, arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
}

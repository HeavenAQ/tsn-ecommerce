package db

import (
	"context"
	"testing"
	"tsn-ecommerce/utils"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	// create a user with random data
	arg := CreateUserParams{
		Email:      utils.RandomNumberString(6) + "@gmail.com",
		Phone:      utils.RandomNumberString(10),
		Password:   utils.RandomAlphabetString(6),
		FirstName:  utils.RandomUserName(),
		LastName:   utils.RandomUserName(),
		LanguagePk: 1,
		Address:    utils.RandomAlphabetString(10),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	// check if there is an error
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// check every field is configured correctly
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.LanguagePk, user.LanguagePk)
	require.Equal(t, arg.Address, user.Address)
	require.NotZero(t, user.Pk)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.LastLogin)
	return user
}

func checkSameUser(t *testing.T, user1, user2 User) {
	require.Equal(t, user1.Pk, user2.Pk)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Phone, user2.Phone)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.LanguagePk, user2.LanguagePk)
	require.Equal(t, user1.Address, user2.Address)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, 0)
	require.WithinDuration(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time, 0)
	require.WithinDuration(t, user1.LastLogin.Time, user2.LastLogin.Time, 0)
}

// Test user creation
func TestCreateUser(t *testing.T) {
	user := createRandomUser(t)
	testQueries.DeleteUser(context.Background(), user.Pk)
}

// Test user retrieval
func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Pk)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	checkSameUser(t, user1, user2)

	// clean up
	testQueries.DeleteUser(context.Background(), user1.Pk)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	checkSameUser(t, user1, user2)

	// clean up
	testQueries.DeleteUser(context.Background(), user1.Pk)
}

func TestGetUserByPhone(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByPhone(context.Background(), user1.Phone)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	checkSameUser(t, user1, user2)

	// clean up
	testQueries.DeleteUser(context.Background(), user1.Pk)
}

func TestListUsers(t *testing.T) {
	num := 10
	orders := make([]User, num)
	for i := 0; i < num; i++ {
		orders[i] = createRandomUser(t)
	}

	// test list users with limit and offset
	args := ListUsersParams{
		Limit:  int32(num),
		Offset: 0,
	}

	// ensure no errors
	users, err := testQueries.ListUsers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, users, 10)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

	// clean up
	for _, user := range orders {
		testQueries.DeleteUser(context.Background(), user.Pk)
	}
}

// Test user update
func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUserParams{
		Pk:         user1.Pk,
		Password:   utils.RandomAlphabetString(6),
		Email:      utils.RandomNumberString(6) + "@gmail.com",
		Phone:      utils.RandomNumberString(10),
		FirstName:  utils.RandomUserName(),
		LastName:   utils.RandomUserName(),
		LanguagePk: 1,
		Address:    utils.RandomAlphabetString(10),
	}

	// update user and check for errors
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	// check if the user is updated correctly
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.Phone, user2.Phone)
	require.Equal(t, arg.FirstName, user2.FirstName)
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, arg.LanguagePk, user2.LanguagePk)
	require.Equal(t, arg.Address, user2.Address)
	require.Equal(t, user1.Pk, user2.Pk)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, user1.CreatedAt.Time, user2.CreatedAt.Time)
	require.NotEqual(t, user1.UpdatedAt.Time, user2.UpdatedAt.Time)
	require.NotEqual(t, user1.LastLogin.Time, user2.LastLogin.Time)

	// clean up
	testQueries.DeleteUser(context.Background(), user1.Pk)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.Pk)
	require.NoError(t, err)

	_, err = testQueries.GetUser(context.Background(), user.Pk)
	require.Error(t, err)
}

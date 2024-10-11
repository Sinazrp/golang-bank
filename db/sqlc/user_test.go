package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomUser() (User, error, CreateUserParams) {
	arg := RandomUser()

	user, err := testQueries.CreateUser(context.Background(), arg)
	return user, err, arg

}
func TestCreateUser(t *testing.T) {
	user, err, arg := CreateRandomUser()
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

}
func TestGetUser(t *testing.T) {
	user1, _, _ := CreateRandomUser()
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, 0)

	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, 0)
}

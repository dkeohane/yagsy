package postgres_test

import (
	"github.com/dkeohane/yagsy/app/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	//"log"
	"testing"
)

func TestCreate(t *testing.T) {
	toCreate := &models.User{}
	toCreate.Email = "yagsyuser@yagsy.com"
	toCreate.Username = "YagsyUser123"
	toCreate.Password = "yagsypass123"

	user, err := store.Create(toCreate)
	require.NoError(t, err)
	assert.NotEqual(t, 0, user.ID)
	assert.Equal(t, toCreate.Email, user.Email)
	assert.Equal(t, toCreate.Username, user.Username)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)

	//	user, err = store.Create(toCreate)

	/*
			account, err = store.Create("authn@keratin.tech", []byte("password"))
			if account != nil {
				assert.NotEqual(t, nil, account)
			}
			if !data.IsUniquenessError(err) {
				t.Errorf("expected uniqueness error, got %T %v", err, err)
			}

		// Assert that db connections are released to pool
		assert.Equal(t, 1, getOpenConnectionCount(store))
	*/
}

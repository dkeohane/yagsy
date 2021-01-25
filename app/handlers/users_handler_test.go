package handlers_test

import (
	"encoding/json"

	"github.com/dkeohane/yagsy/app/data"
	//"github.com/dkeohane/yagsy/app/dbs"
	//"github.com/dkeohane/yagsy/app/handlers"
	"github.com/dkeohane/yagsy/app/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang.org/x/crypto/bcrypt"
	"net/http"
	//"os"
	//"log"
	"testing"
)

/*func TestMain(m *testing.M) {

	// Setup a mocked DB
	mockedUserStore, err := dbs.NewUserStore("mocked", nil)
	if err != nil {
		//t.Error("Error calling NewUserStore, err: " + err.Error())
		log.Fatal(err)
	}
	MyUserHandler = handlers.UserHandler{DB: mockedUserStore}

	m.Run()
}*/

func TestCreateUser(t *testing.T) {

	var Email string = "yagsyuser@yagsy.com"
	var Username string = "YagsyUser123"
	var Password string = "yagsypass123"

	body := string(`
		{
			"username":"YagsyUser123",
			"email": "yagsyuser@yagsy.com",
			"password": "yagsypass123"
		}
	`)

	test := utils.GenerateHandleTester(t, MyUserHandler.CreateUser)
	w := test(http.MethodPost, body)

	if w.Code != http.StatusOK {
		t.Errorf("CreateUser didn't return a created user %v", http.StatusOK)
		return
	}

	user := &models.User{}
	err := json.NewDecoder(w.Body).Decode(user)

	if err != nil {
		t.Errorf("%v", err)
	}

	require.NoError(t, err)
	assert.NotEqual(t, 0, user.ID)
	assert.Equal(t, Email, user.Email)
	assert.Equal(t, Username, user.Username)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))

	if err != nil {
		t.Errorf("%v", err)
	}

	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}

/*
func TestGetUser(t *testing.T) {

	var Email string = "yagsyuser@yagsy.com"
	var Username string = "YagsyUser123"
	var Password string = "yagsypass123"

	// Setup a mocked DB
	mockedUserStore, err := dbs.NewUserStore("mocked", nil)
	if err != nil {
		t.Error("Error calling NewUserStore, err: " + err.Error())
	}
	MyUserHandler := &handlers.UserHandler{DB: mockedUserStore}

	body := string(`
		{
			"username":"YagsyUser123",
			"email": "yagsyuser@yagsy.com",
			"password": "Yagsypass123"
		}
	`)

	test := utils.GenerateHandleTester(t, MyUserHandler.CreateUser)
	w := test(http.MethodPost, body)

	if w.Code != http.StatusOK {
		t.Errorf("CreateUser didn't return a created user %v", http.StatusOK)
		return
	}

	user := &models.User{}
	err = json.NewDecoder(w.Body).Decode(user)

	if err != nil {
		t.Errorf("%v", err)
	}

	require.NoError(t, err)
	assert.NotEqual(t, 0, user.ID)
	assert.Equal(t, Email, user.Email)
	assert.Equal(t, Username, user.Username)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))

	if err != nil {
		t.Errorf("%v", err)
	}

	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}
*/

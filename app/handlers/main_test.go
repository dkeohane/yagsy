package handlers_test

import (
	"github.com/dkeohane/yagsy/app/dbs"
	"github.com/dkeohane/yagsy/app/handlers"

	"log"
	"os"
	"testing"
)

var MyUserHandler handlers.UserHandler

func TestMain(m *testing.M) {
	mockedUserStore, err := dbs.NewUserStore("mocked", nil)
	if err != nil {
		log.Fatal(err)
	}
	MyUserHandler = handlers.UserHandler{DB: mockedUserStore}

	os.Exit(m.Run())
}

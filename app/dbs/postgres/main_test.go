package postgres_test

import (
	"github.com/dkeohane/yagsy/app/dbs/postgres"
	"github.com/dkeohane/yagsy/config"
	"log"
	"os"
	"testing"
)

var store postgres.UserStore

func TestMain(m *testing.M) {
	var testConfig config.Config
	err := testConfig.LoadConfig("../../../config/config_test.yaml", true)

	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewDB(&testConfig)
	store = postgres.UserStore{Session: db}
	store.CreateTables()

	os.Exit(m.Run())
}

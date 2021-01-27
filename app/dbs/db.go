package dbs

import (
	"database/sql"
	//"fmt"
	"github.com/dkeohane/yagsy/app/dbs/postgres"
	"github.com/dkeohane/yagsy/app/utils/config"
)

// NewDB allows us to have more than one DB provider
func NewDB(config *config.Config) (*sql.DB, error) {

	return postgres.NewDB(config)
	/*
		switch url {
		case "postgresql", "postgres":
			return postgres.NewDB(url)
		default:
			return nil, fmt.Errorf("Unsupported database: %s", url)
		}*/
}

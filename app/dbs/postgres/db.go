package postgres

import (
	"database/sql"
	"fmt"
	"github.com/dkeohane/yagsy/app/utils/config"
	_ "github.com/lib/pq"
)

func NewDB(config *config.Config) (*sql.DB, error) {

	// DBInfo string, should be moved out of here at a later point.
	dbInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name)

	client, err := sql.Open("postgres", dbInfo)
	CheckError(err)
	err = client.Ping()
	CheckError(err)
	return client, nil
}

package dbs

import (
	"database/sql"
	"github.com/dkeohane/yagsy/app/data"
	"github.com/dkeohane/yagsy/app/dbs/mock"
	"github.com/dkeohane/yagsy/app/dbs/postgres"
	"github.com/google/uuid"
)

type UserStore interface {
	Create(user *models.User) (*models.User, error)
	GetUser(userID uuid.UUID) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUsers(filterkey string, filterlist []string) ([]models.User, error)
	UpdateUser(userID uuid.UUID, user *models.User) (*models.UpdateResult, error)
	DeleteUser(userID uuid.UUID) error
	CreateTables() error
}

func NewUserStore(dbType string, db *sql.DB) (UserStore, error) {
	// Here depending on the DB type we would return a different
	// User Store as well. Currently, only postgres is supported.
	if dbType == "postgres" {
		return &postgres.UserStore{Session: db}, nil
	} else {
		return &mocked.UserStore{}, nil
	}

}

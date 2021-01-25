package mocked

import (
	"github.com/dkeohane/yagsy/app/data"
	"github.com/google/uuid"
	"time"
)

type UserStore struct{}

func (m *UserStore) Create(user *models.User) (*models.User, error) {

	user.ID = uuid.New()
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	return user, nil
}

func (m *UserStore) GetUser(userID uuid.UUID) (*models.User, error) {
	user := models.User{}
	return &user, nil
}

func (m *UserStore) GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}
	return &user, nil
}

func (m *UserStore) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	return &user, nil
}
func (m *UserStore) GetUsers(filterkey string, filterlist []string) ([]models.User, error) {
	users := []models.User{}
	return users, nil
}

func (m *UserStore) UpdateUser(userID uuid.UUID, user *models.User) (*models.UpdateResult, error) {
	return &models.UpdateResult{}, nil
}

func (m *UserStore) DeleteUser(userID uuid.UUID) error {
	return nil
}

func (m *UserStore) CreateTables() error {
	return nil
}

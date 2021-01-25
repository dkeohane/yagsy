package postgres

import (
	"context"
	"database/sql"
	"github.com/dkeohane/yagsy/app/data"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
	"time"
)

type UserStore struct {
	Session *sql.DB
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (m *UserStore) Create(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	queryString := `
		INSERT INTO Users(
			id,
			email,
			password,
			username,
			tokenhash,
            isverified,
            createdat,
            updatedat
		) VALUES ($1, $2, $3, $4, $5, $6, $7::timestamp, $8::timestamp);
	`

	user.ID = uuid.New()
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	_, err := m.Session.ExecContext(
		ctx,
		queryString,
		user.ID,
		user.Email,
		user.Password,
		user.Username,
		user.TokenHash,
		user.IsVerified,
		user.CreatedAt,
		user.UpdatedAt,
	)
	CheckError(err)
	return user, err
}

func (m *UserStore) GetUser(userID uuid.UUID) (*models.User, error) {
	queryString := `SELECT * FROM Users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	user := &models.User{}

	m.Session.QueryRowContext(
		ctx,
		queryString,
		userID,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.TokenHash,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, nil
}

func (m *UserStore) GetUserByUsername(username string) (*models.User, error) {
	queryString := `SELECT * FROM Users WHERE username = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	user := &models.User{}

	m.Session.QueryRowContext(
		ctx,
		queryString,
		username,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.TokenHash,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, nil
}

func (m *UserStore) GetUserByEmail(email string) (*models.User, error) {
	queryString := `SELECT * FROM Users WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	user := &models.User{}

	m.Session.QueryRowContext(
		ctx,
		queryString,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.TokenHash,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, nil
}

func (m *UserStore) GetUsers(filterkey string, filterlist []string) ([]models.User, error) {

	users := []models.User{}
	query := `
	SELECT
		id,
		email,
		password,
		username,
		tokenhash,
		isverified,
		createdat,
		updatedat
	FROM users
	WHERE ` + filterkey + " = ANY($1) ORDER BY id DESC"
	param := "{" + strings.Join(filterlist, ",") + "}"

	//  LIMIT $1 OFFSET $2

	/*
			func getProducts(db *sql.DB, start, count int) ([]product, error) {
		    rows, err := db.Query(
		        "SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
		        count, start)

		    if err != nil {
		        return nil, err
		    }

		    defer rows.Close()

		    products := []product{}

		    for rows.Next() {
		        var p product
		        if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
		            return nil, err
		        }
		        products = append(products, p)
		    }

		    return products, nil
		}
	*/

	log.Println(query)
	log.Println(param)

	rows, err := m.Session.Query(query, param)

	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.Username,
			&user.TokenHash,
			&user.IsVerified,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		CheckError(err)
		users = append(users, user)
		// repos.Repositories = append(repos.Repositories, repo)
		// ADD USER TO USERS LIST
	}

	err = rows.Err()
	CheckError(err)

	return users, nil
}

func (m *UserStore) UpdateUser(userID uuid.UUID, user *models.User) (*models.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	user.UpdatedAt = time.Now().UTC()

	valueArgs := []interface{}{}
	queryString := "UPDATE Users SET "

	if user.Username != "" {
		queryString += " username = $" + strconv.Itoa(len(valueArgs)+1) + ","
		valueArgs = append(valueArgs, user.Username)
	}

	if user.Email != "" {
		queryString += " email = $" + strconv.Itoa(len(valueArgs)+1) + ","
		valueArgs = append(valueArgs, user.Email)
	}

	if user.Password != "" {
		queryString += " password = $" + strconv.Itoa(len(valueArgs)+1) + ","
		valueArgs = append(valueArgs, user.Password)
	}

	queryString += " updatedat = $" + strconv.Itoa(len(valueArgs)+1)
	valueArgs = append(valueArgs, user.UpdatedAt)

	queryString += " WHERE id = $" + strconv.Itoa(len(valueArgs)+1)
	valueArgs = append(valueArgs, userID)

	queryString += " RETURNING username, email, password"

	var username, email, password string
	m.Session.QueryRowContext(ctx, queryString, valueArgs...).Scan(&username, &email, &password)

	result := &models.UpdateResult{userID, username, email, password}

	return result, nil
}

func (m *UserStore) DeleteUser(userID uuid.UUID) error {
	queryString := `DELETE FROM Users WHERE id = $1`
	_, err := m.Session.Exec(queryString, userID)
	CheckError(err)
	return err
}

func (m *UserStore) CreateTables() error {
	log.Println("Creating DB Tables")

	queryString := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	_, err := m.Session.Exec(queryString)
	CheckError(err)

	queryString = `
		CREATE TABLE IF NOT EXISTS Users (
			id uuid DEFAULT uuid_generate_v4 () NOT NULL UNIQUE,
			email varchar(111) NOT NULL UNIQUE,
			password varchar(111) NOT NULL,
			username varchar(111) NOT NULL,
			tokenhash varchar(111),
			isverified boolean,
			createdat timestamp,
			updatedat timestamp,
			PRIMARY KEY (id)
		);
	`
	_, err = m.Session.Exec(queryString)
	CheckError(err)
	return err
}

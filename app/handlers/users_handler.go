package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dkeohane/yagsy/app/data"
	"github.com/dkeohane/yagsy/app/dbs"
	"github.com/dkeohane/yagsy/app/utils/token"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type UserHandler struct {
	JwtSecret string
	DB        dbs.UserStore
}

// ApplyRoutes applies the User routes to the mux Router
func (uh *UserHandler) ApplyRoutes(router *mux.Router, authHandler *AuthHandler) {

	log.Println("Applying User Routes")

	// Create User
	router.HandleFunc("/v1.0/users", uh.CreateUser).Methods(http.MethodPost)

	// Get Users, can supply filterKey
	router.Handle("/v1.0/users", authHandler.authHandler(uh.GetUsers)).Methods(http.MethodGet)

	// Get Specific User Data
	router.HandleFunc("/v1.0/users/{id}", uh.GetUser).Methods(http.MethodGet)

	// Update User
	router.HandleFunc("/v1.0/users/{id}", uh.UpdateUser).Methods(http.MethodPut)

	// Delete User
	router.HandleFunc("/v1.0/users/{id}", uh.DeleteUser).Methods(http.MethodDelete)

	// Login User
	router.HandleFunc("/v1.0/login", uh.Login).Methods(http.MethodPost)
}

type ErrorResponse struct {
	Err string
}

// CreateUser creates a new user with the provided data from Request.
func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Error decoding request payload")

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))

		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)
	user, err = uh.DB.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(user)
}

// GetUser gets a specific user from the DB. This user is
// identified by their id.
func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := uuid.Parse(params["id"])
	CheckError(err)
	user, err := uh.DB.GetUser(userID)
	CheckError(err)
	json.NewEncoder(w).Encode(&user)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetUsers encodes a list of Users found in the DB into the provided
// ResponseWriter. The returned list is filtered on filterKey. Currently
// supported filter keys:
//   - username
func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	parsedQuery, err := url.ParseQuery(r.URL.RawQuery)
	CheckError(err)

	filterKeyGroup := parsedQuery["filterKey"]
	if len(filterKeyGroup) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - filterKey should be a string value"))
		return
	}

	filterKey := filterKeyGroup[0]
	if !filterKeyIsValid(filterKey) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - filterKey is invalid"))
		return
	}

	filterData := parsedQuery["filterData"] // params["filterData"]

	users, err := uh.DB.GetUsers(filterKey, filterData)
	CheckError(err)
	json.NewEncoder(w).Encode(&users)
}

func filterKeyIsValid(filterKey string) bool {
	fKey := strings.ToLower(filterKey)
	if fKey != "username" {
		return false
	}
	return true
}

// UpdateUser updates the specified user
func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userID, err := uuid.Parse(params["id"])

	CheckError(err)

	user := &models.User{}
	err = json.NewDecoder(r.Body).Decode(user)

	CheckError(err)

	user.ID = userID
	updateResult, err := uh.DB.UpdateUser(userID, user)
	CheckError(err)
	json.NewEncoder(w).Encode(&updateResult)
}

// DeleteUser deletes the specified user
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := uuid.Parse(params["id"])
	CheckError(err)
	err = uh.DB.DeleteUser(userID)
	CheckError(err)
	json.NewEncoder(w).Encode("User deleted")
}

// CreateUser creates a new user with the provided data from Request.
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)
	retrievedUser, err := uh.DB.GetUserByUsername(user.Username)

	err = bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(user.Password))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	// OK user is correct, let's create a token!
	authToken, err := token.New(
		retrievedUser.ID,
		"github.com/dkeohane/yagsy").SignHS256([]byte(uh.JwtSecret))

	CheckError(err)

	authRetMap := map[string]string{
		"Authorization": authToken,
	}

	authRetJSON, err := json.Marshal(authRetMap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(authRetJSON)
	}

}

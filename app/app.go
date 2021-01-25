// app.go

package app

import (
	"database/sql"
	"github.com/dkeohane/yagsy/app/dbs"
	"github.com/dkeohane/yagsy/app/handlers"
	"github.com/dkeohane/yagsy/config"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App struct {
	Router      *mux.Router
	DB          *sql.DB
	Config      *config.Config
	UserHandler *handlers.UserHandler
	AuthHandler *handlers.AuthHandler
}

func (app *App) Initialize() {

	// Instantiate the gorilla/mux router.
	app.Router = mux.NewRouter()

	app.Config = &config.Config{}
	err := app.Config.LoadConfig("/config/config.yaml", false)
	if err != nil {
		log.Fatal(err)
	}

	// Create DB
	app.DB, err = dbs.NewDB(app.Config)
	if err != nil {
		log.Fatal(err)
	}

	// Create UserStore matching the DB we've created
	userStore, err := dbs.NewUserStore("postgres", app.DB)

	// Create Tables for the DB
	userStore.CreateTables()

	// Give the UserHandler the newly created UserStore - it shouldn't
	// care what DB is used, as they implement a common interface
	app.UserHandler = &handlers.UserHandler{
		DB:        userStore,
		JwtSecret: app.Config.Auth.JwtSecret,
	}
	app.AuthHandler = &handlers.AuthHandler{
		JwtSecret: app.Config.Auth.JwtSecret,
	}

	// Apply User Routes to serve.
	app.UserHandler.ApplyRoutes(app.Router, app.AuthHandler)
}

func (app *App) Run() {
	http.ListenAndServe(":8080", app.Router)
}

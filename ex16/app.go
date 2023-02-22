package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}

func (a *App) Initialize(user, passwd, dbname string) {}

func (a *App) Run(addr string) {}

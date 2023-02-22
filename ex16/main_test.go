package main_test

import (
	"os"
	"tesing"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"encoding/json"
	"bytes"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.exec(tableCreationQuery); err != nil {
		log.fatal(errr)
	}
}

func clearTable() {
	a.DB.Exec("Delete from products")
	a.DB.Exec("Alter SEQUENCE products_id_seq RESTART wWITH 1")

}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`
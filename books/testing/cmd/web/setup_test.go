package main

import (
	"log"
	"os"
	"testing"

	"github.com/MinhNHHH/testing/pkg/db"
)

var app application

func TestMain(m *testing.M) {
	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"
	app.Session = getSession()
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	app.DB = db.PostgresConn{DB: conn}

	pathToTemplates = "/home/minh/Desktop/learn_cs/books/testing/cmd/template/"
	os.Exit(m.Run())
}

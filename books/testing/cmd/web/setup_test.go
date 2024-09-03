package main

import (
	"os"
	"testing"

	"github.com/MinhNHHH/testing/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	app.Session = getSession()
	app.DB = &dbrepo.TestDBRepo{}

	pathToTemplates = ".././template/"
	os.Exit(m.Run())
}

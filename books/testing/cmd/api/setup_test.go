package main

import (
	"os"
	"testing"

	"github.com/MinhNHHH/testing/pkg/repository/dbrepo"
)

var app application

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "abcde"

	os.Exit(m.Run())
}

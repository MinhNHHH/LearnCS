package main

import (
	"os"
	"testing"
)

var app application

func TestMain(m *testing.M) {
	app.Session = getSession()
	pathToTemplates = "/home/minh/Desktop/learn_cs/books/testing/cmd/template/"
	os.Exit(m.Run())
}

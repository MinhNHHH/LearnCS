package main

import (
	"os"
	"testing"

	"github.com/MinhNHHH/testing/pkg/repository/dbrepo"
)

var app application
var expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXVkIjoiZXhhbXBsZS5jb20iLCJleHAiOjE3MjUyNDU0NTUsImlzcyI6ImV4YW1wbGUuY29tIiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMSJ9.L0znjDqQI2ApiKEl3RyO6ou0rAEcS1fuZZ1emH3aGXU"

func TestMain(m *testing.M) {
	app.DB = &dbrepo.TestDBRepo{}
	app.Domain = "example.com"
	app.JWTSecret = "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160"

	os.Exit(m.Run())
}

package app

import (
	"github.com/MinhNHHH/go_dev/pkg/db"
)

type Application struct {
	JWTSecret string
	DB        *db.Database
}

func NewApplication(JWTSecret, DSN string) *Application {
	return &Application{
		JWTSecret: JWTSecret,
		DB:        db.Init(DSN),
	}
}

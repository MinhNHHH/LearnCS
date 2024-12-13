package main

import (
	"flag"
	"fmt"

	ap "github.com/MinhNHHH/go_dev/pkg/app"
)

const port = 8080

type application struct {
	DSN       string
	JWTSecret string
}

func main() {
	var app application
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable", "Posgtres connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "signing secret")
	flag.Parse()

	application := ap.NewApplication(app.JWTSecret, app.DSN)
	routes := application.Routes()
	routes.Run(fmt.Sprintf(":%d", port))
}

package main

import (
	"flag"
	"fmt"

	ap "github.com/MinhNHHH/go_dev/pkg/app"
)

const port = 8080

type application struct {
	DSN       string
	Domain    string
	JWTSecret string
}

func main() {
	var app application
	flag.StringVar(&app.Domain, "domain", "example.com", "Domain for application, e.g. company.com")
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Posgtres connection")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "signing secret")
	flag.Parse()

	application := ap.NewApplication(app.JWTSecret, app.DSN)
	r := application.Routes()
	r.Run(fmt.Sprintf(":%d", port))
}

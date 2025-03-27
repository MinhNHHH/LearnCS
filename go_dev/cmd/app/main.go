package main

import (
	"fmt"

	do "github.com/MinhNHHH/go_dev/pkg/documents"
)

const port = 8080

type application struct {
	DSN       string
	JWTSecret string
}

func main() {
	// var app application
	// flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable", "Posgtres connection")
	// flag.StringVar(&app.JWTSecret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "signing secret")
	// flag.Parse()
	//
	// application := ap.NewApplication(app.JWTSecret, app.DSN)
	// routes := application.Routes()
	// routes.Run(fmt.Sprintf(":%d", port))
	seedURL := "https://www.linkedin.com/jobs/search/?currentJobId=4122800813&geoId=104195383&keywords=golang&origin=JOB_SEARCH_PAGE_SEARCH_BUTTON"
	crawler := do.NewCrawler()
	resp, err := crawler.FetchPage(seedURL)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		return
	}

	job, err := crawler.JobParser.Parser(resp, seedURL)

	if err != nil {
		fmt.Println("Error extracting links:", err)
		return
	}
	fmt.Println(job)
}

package ex410

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Id        int
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	// in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func ex410() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// now := time.Now()
	// pastDay := now.AddDate(0, 0, -1)
	// pastMonth := now.AddDate(0, -1, 0)
	// pastYear := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		// if item.CreatedAt.After(pastYear) || item.CreatedAt.After(pastMonth) || item.CreatedAt.After(pastDay) {
		// 	fmt.Printf("#%-5d %9.9s %.55s %d\n",
		// 		item.Number, item.User.Login, item.Title, item.Id)
		// }
		fmt.Println(item)
	}
}

const IssuesURL = "https://api.github.com/search/issues"

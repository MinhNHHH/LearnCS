package exercises

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"text/template"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
	columns    []Comparation
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type Comparation func(a, b *Issue) int

func SortByNumber(a, b *Issue) int {
	switch {
	case a.Number < b.Number:
		return lt
	case a.Number > b.Number:
		return gt
	case a.Number == b.Number:
		return eq
	}
	return -1
}

func SortByTitle(a, b *Issue) int {
	switch {
	case a.Title < b.Title:
		return lt
	case a.Title > b.Title:
		return gt
	case a.Title == b.Title:
		return eq
	}
	return -1
}

func SortByState(a, b *Issue) int {
	switch {
	case a.State < b.State:
		return lt
	case a.State > b.State:
		return gt
	case a.State == b.State:
		return eq
	}
	return -1
}

func (s IssuesSearchResult) Len() int { return len(s.Items) }
func (s IssuesSearchResult) Less(x, y int) bool {
	for _, fsort := range s.columns {
		cmp := fsort(s.Items[x], s.Items[y])
		switch cmp {
		case lt:
			return true
		case gt:
			return false
		case eq:
			return false
		}
	}
	return false
}
func (s IssuesSearchResult) Swap(x, y int) {
	s.Items[x], s.Items[y] = s.Items[y], s.Items[x]
}
func (s *IssuesSearchResult) Select(cmp Comparation) {
	s.columns = append([]Comparation{cmp}, s.columns...)
}

func NewIssuesSearchResults(issues []*Issue) *IssuesSearchResult {
	return &IssuesSearchResult{Items: issues}
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const IssuesURL = "https://api.github.com/search/issues"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	terms := r.URL.Query()["q"]
	sortKeys := strings.Split(r.URL.Query()["sort"][0], ",")
	if len(terms) == 0 {
		http.Error(w, "No search terms provided", http.StatusBadRequest)
		return
	}
	result, err := SearchIssues(terms)
	for _, sortKey := range sortKeys {
		switch strings.ToLower(sortKey) {
		case "number":
			result.Select(SortByNumber)
		case "state":
			result.Select(SortByState)
		case "title":
			result.Select(SortByTitle)
		}
	}

	sort.Sort(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Search query failed: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := issueList.Execute(w, result); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render template: %s", err), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/search", issueHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

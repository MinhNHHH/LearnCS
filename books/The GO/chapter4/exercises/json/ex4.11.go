package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

type IssueMapSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      map[int]*Issue
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

const IssuesURL = "https://api.github.com/search/issues"

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
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

func readFileExist(filePath string) *IssueMapSearchResult {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil
	}

	var issuess IssueMapSearchResult
	if err := json.Unmarshal(jsonData, &issuess); err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v\n", err)
		return nil
	}
	return &issuess
}

func fetchNew(data *IssuesSearchResult) string {

	var mapSearchResult IssueMapSearchResult
	mapSearchResult.TotalCount = data.TotalCount
	mapSearchResult.Items = make(map[int]*Issue)

	for _, item := range data.Items {
		mapSearchResult.Items[item.Id] = item
	}

	jsonData, jsonErr := json.MarshalIndent(mapSearchResult, "", " ")
	if jsonErr != nil {
		return fmt.Sprintf("Failed to marshal JSON: %v\n", jsonErr)
	}

	// Write JSON data to a file
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	return fmt.Sprintf("%d", data.TotalCount)
}

func fetchIssues(terms []string) string {
	issues, err := SearchIssues(terms)
	if err != nil {
		return "Error when collect issues from Github"
	}
	total := "Fetch total issues: "
	if issues.TotalCount > 0 {
		if _, err := os.Stat("data.json"); os.IsNotExist(err) {
			total = total + fetchNew(issues)
		} else {
			existIssues := readFileExist("data.json")
			countIssues := 0

			for _, item := range issues.Items {
				if _, ok := existIssues.Items[item.Id]; !ok {
					existIssues.Items[item.Id] = item
					countIssues++
				}
			}
			existIssues.TotalCount += countIssues

			// Marshal the updated struct back to JSON
			updatedJSON, err := json.MarshalIndent(existIssues, "", "    ")
			if err != nil {
				return fmt.Sprintf("Failed to marshal JSON: %v\n", err)
			}

			// Write the JSON data back to the file
			if err := ioutil.WriteFile("data.json", updatedJSON, 0644); err != nil {
				return fmt.Sprintf("Failed to write file: %v\n", err)
			}
			total = total + fmt.Sprintf("%d", countIssues)
		}
	}

	return total
}

func DisplayIssues() {
	issues := readFileExist("data.json")
	fmt.Printf("%s\t\t%12s\t\t%13s\t\t%.55s\n", "ID", "USER", "NUMBER", "TITLE")
	for _, value := range issues.Items {
		fmt.Printf("%d\t\t%9.9s\t\t%d\t\t%.55s\n", value.Id, value.User.Login, value.Number, value.Title)
	}
}

func deleteIssuesById(id int) string {
	issues := readFileExist("data.json")
	if _, ok := issues.Items[id]; ok {
		delete(issues.Items, id)
	} else {
		return fmt.Sprintf("Cannot find issue has id %d", id)
	}
	// Marshal the updated struct back to JSON
	updatedJSON, err := json.MarshalIndent(issues, "", "    ")
	if err != nil {
		return fmt.Sprintf("Failed to marshal JSON: %v\n", err)
	}

	// Write the JSON data back to the file
	if err := ioutil.WriteFile("data.json", updatedJSON, 0644); err != nil {
		return fmt.Sprintf("Failed to write file: %v\n", err)
	}
	return fmt.Sprintf("Delete issue has id %d success", id)
}

func main() {
	var fetch bool
	var read bool
	var delete int
	flag.BoolVar(&fetch, "fetch", false, "Fetch issues from github")
	flag.IntVar(&delete, "delete", 0, "Delete issues by id")
	flag.BoolVar(&read, "read", false, "Read issues from github")

	flag.Parse()

	if fetch {
		fmt.Println(fetchIssues(os.Args[2:]))
	} else if read {
		DisplayIssues()
	} else if delete != 0 {
		fmt.Println(deleteIssuesById(delete))
	}
}

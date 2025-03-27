package documents

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Attribute struct {
	Namespace, Key, Val string
}

type Node struct {
	Type     string
	Data     string
	Attr     []Attribute
	Children []Node
}

type Queue[T any] struct {
	Items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var item T
	if len(q.Items) == 0 {
		return item, fmt.Errorf("queue is empty")
	}

	item = q.Items[0]
	q.Items = q.Items[1:]

	return item, nil
}

func (q *Queue[T]) IsEmpty() bool {
	if len(q.Items) == 0 {
		return true
	}
	return false
}

func (q *Queue[T]) Size() int {
	return len(q.Items)
}

type UrlFrontier struct {
	Queue   *Queue[string]
	Visited map[string]bool
	mux     sync.Mutex
}

type SeedURL struct{}

type Crawler struct {
	UrlFrontier *UrlFrontier
	SeedUrl     SeedURL
	JobParser   *JobParser
}
type JobPosting struct {
	URL          string
	Title        string
	Company      string
	Location     string
	Description  string
	Salary       string
	Requirements []string
	PostedDate   string
}

type JobParser struct {
	Job *JobPosting
}

func NewJobParser() *JobParser {
	return &JobParser{
		Job: &JobPosting{},
	}
}

func NewCrawler() *Crawler {
	return &Crawler{
		UrlFrontier: &UrlFrontier{
			Visited: map[string]bool{},
			Queue:   &Queue[string]{}},
		SeedUrl:   SeedURL{},
		JobParser: NewJobParser(),
	}
}

func (c *Crawler) FetchPage(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	return resp, nil
}

func isValidURL(rawURL string) bool {
	parsedURL, err := url.Parse(rawURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}

func getContent(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if isValidURL(attr.Val) {
				// fmt.Println(attr.Val)
			}
		}
	}

	if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			fmt.Println("dddd", attr.Val)
		}
	}

	// if n.Type == html.TextNode {
	// 	// return strings.TrimSpace(n.Data)
	// }

	// var text []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getContent(c)
		// if content := getTextContent(c); content != "" {
		// 	text = append(text, content)
		// }
	}
	// return strings.Join(text, " ")
}

func (c *Crawler) ExtractLinks(resp *http.Response) error {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	var extract func(n *html.Node)
	extract = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if isValidURL(attr.Val) && !c.UrlFrontier.Visited[attr.Val] {
					c.UrlFrontier.Queue.Enqueue(attr.Val)
					c.UrlFrontier.Visited[attr.Val] = true
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	extract(doc)
	return nil
}

func (jp *JobParser) extractJobInfo(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, attr := range n.Attr {
				// Get company name
				if attr.Key == "class" && strings.Contains(attr.Val, "hidden-nested-link") {
					if n.FirstChild != nil {
						jp.Job.Company = strings.TrimSpace(n.FirstChild.Data)
					}
				} else {
					continue
				}

				if isValidURL(attr.Val) {
					jp.Job.URL = attr.Val
				}
			}
		case "span":
			// Get location
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "job-search-card__location") {
					if n.FirstChild != nil {
						jp.Job.Location = strings.TrimSpace(n.FirstChild.Data)
					}
				}
			}
		case "h3":
			// Get title
			for _, attr := range n.Attr {
				if attr.Key == "class" && strings.Contains(attr.Val, "base-search-card__title") {
					jp.Job.Title = strings.TrimSpace(n.FirstChild.Data)
				}
			}
		case "time":
			// Get posted date
			for _, attr := range n.Attr {
				if attr.Key == "datetime" {
					jp.Job.PostedDate = strings.TrimSpace(attr.Val)
				}
			}
		case "div":
			for _, attr := range n.Attr {
				if n.FirstChild != nil {
					fmt.Println(strings.TrimSpace(n.FirstChild.Data), "===========", attr.Key, "================", attr.Val)
					// jp.Job.Description = strings.TrimSpace(n.FirstChild.Data)
				}

				if attr.Key == "class" && strings.Contains(attr.Val, "job-description") {
					// Get description
				}
			}
		}

	}

	// Continue traversing the DOM tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		jp.extractJobInfo(c)
	}
}
func (jp *JobParser) Parser(resp *http.Response, url string) (*JobPosting, error) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	jp.extractJobInfo(doc)
	return jp.Job, nil
}

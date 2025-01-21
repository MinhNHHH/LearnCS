package documents

import (
	"fmt"
	"net/http"
	"net/url"
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
}

func NewCrawler() *Crawler {
	return &Crawler{
		UrlFrontier: &UrlFrontier{
			Visited: map[string]bool{},
			Queue:   &Queue[string]{}},
		SeedUrl: SeedURL{},
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

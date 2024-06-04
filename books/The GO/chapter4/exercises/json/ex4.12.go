package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

const (
	xkcdURL  = "https://xkcd.com"
	filePath = "comics/store.json"
)

// XKCDSearchResult represents a single comic's metadata
type XKCDSearchResult struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

type Comic int

// Index represents an index of terms to comics
type Index struct {
	Index map[string]map[Comic]bool
	mu    sync.RWMutex
}

// ComicIndex holds the complete index and the comics data
type ComicIndex struct {
	Index  map[string]map[Comic]bool
	Comics map[int]XKCDSearchResult
}

func main() {
	var fetch string
	var search string
	flag.StringVar(&fetch, "fetch", "", "Fetch comics based on ids")
	flag.StringVar(&search, "search", "", "Search string")
	flag.Parse()

	if fetch != "" {
		terms := strings.Split(fetch, ",")
		comics, total := fetchComics(terms)

		index := newIndex()
		index.buildIndex(comics)

		err := storeComics(ComicIndex{Index: index.Index, Comics: comics})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Total fetch is %d\n", total)
	} else if search != "" {
		comics, err := loadComics(filePath)
		if err != nil {
			log.Fatal(err)
		}
		searchComics(search, *comics)
	}
}

// newIndex initializes a new Index.
func newIndex() *Index {
	return &Index{
		Index: make(map[string]map[Comic]bool),
	}
}

// tokenize splits a document into terms.
func tokenize(content string) []string {
	return strings.Fields(content)
}

// cleanContent removes non-alphanumeric characters from the content.
func cleanContent(content string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	return re.ReplaceAllString(strings.TrimSpace(content), "")
}

// buildIndex builds the index from the provided comics.
func (idx *Index) buildIndex(comics map[int]XKCDSearchResult) {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	for _, comic := range comics {
		cleanedContent := cleanContent(comic.Transcript)
		terms := tokenize(cleanedContent)
		for _, term := range terms {
			normalizedTerm := strings.ToLower(term)
			if idx.Index[normalizedTerm] == nil {
				idx.Index[normalizedTerm] = make(map[Comic]bool)
			}
			idx.Index[normalizedTerm][Comic(comic.Num)] = true
		}
	}
}

// storeComics saves the comics data and index to a JSON file.
func storeComics(data ComicIndex) error {
	var comicsIndex ComicIndex

	if _, err := os.Stat(filePath); err == nil {
		comics, err := loadComics(filePath)
		if err != nil {
			log.Printf("Failed to load existing comics: %v\n", err)
			comicsIndex = data
		} else {
			comicsIndex = *comics
		}
	} else {
		comicsIndex = data
	}

	// Merge new data into existing comics
	for _, comic := range data.Comics {
		comicsIndex.Comics[comic.Num] = comic
	}

	// Update the Index map with new data
	for word, comicsMap := range data.Index {
		if existingComicsMap, ok := comicsIndex.Index[word]; ok {
			for comic, present := range comicsMap {
				existingComicsMap[comic] = present
			}
		} else {
			comicsIndex.Index[word] = comicsMap
		}
	}

	jsonData, err := json.MarshalIndent(comicsIndex, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	err = os.MkdirAll("comics", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

// fetchComics fetches comics based on the given terms.
func fetchComics(terms []string) (map[int]XKCDSearchResult, int) {
	total := 0
	comics := make(map[int]XKCDSearchResult)
	for _, term := range terms {
		data, err := searchComic(term)
		if err != nil {
			log.Fatalf("Failed to fetch data for comic %s: %v", term, err)
		}
		comics[data.Num] = *data
		total++
	}
	return comics, total
}

// searchComics searches for the given term in the comics index.
func searchComics(term string, comicIndex ComicIndex) {
	normalizedTerm := strings.ToLower(term)
	docPositions, found := comicIndex.Index[normalizedTerm]
	if !found {
		return
	}
	for docID := range docPositions {
		comic := comicIndex.Comics[int(docID)]
		fmt.Printf("URL: %s/%d/\n", xkcdURL, comic.Num)
		fmt.Printf("Transcript: %s\n\n", comic.Transcript)
		fmt.Println("==========================================")
	}
}

// searchComic fetches a comic by its ID.
func searchComic(id string) (*XKCDSearchResult, error) {
	url := fmt.Sprintf("%s/%s/info.0.json", xkcdURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comic: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch comic: received status code %d", resp.StatusCode)
	}

	var result XKCDSearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode comic JSON: %v", err)
	}

	return &result, nil
}

// loadComics loads comics from a JSON file.
func loadComics(path string) (*ComicIndex, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var comicsIndex ComicIndex
	err = json.Unmarshal(data, &comicsIndex)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &comicsIndex, nil
}

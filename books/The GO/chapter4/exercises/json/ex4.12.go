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

const xkcdURL = "https://xkcd.com"
const filePath = "comics/store.json"

type XKCDSearchResult struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}
type Comic int

type Index struct {
	Index map[string]map[Comic]bool
	mu    sync.RWMutex
}

type ComicIndex struct {
	Index  map[string]map[Comic]bool
	Comics map[int]XKCDSearchResult
}

func main() {
	var fetch string
	var search string
	flag.StringVar(&fetch, "fetch", "", "Fetch comics base on ids")
	flag.StringVar(&search, "search", "", "Search string")
	flag.Parse()

	if fetch != "" {
		terms := strings.Split(fetch, ",")
		comics, total := Fetch(terms)

		index := NewIndex()
		index.BuildIndex(comics)

		StoreComics(ComicIndex{Index: index.Index, Comics: comics})
		fmt.Printf("Total fetch is %d\n", total)
	} else if search != "" {
		comics, err := loadComics("comics/store.json")
		if err != nil {
			log.Fatal(err)
		}
		Search(search, *comics)
	}
}

// NewIndex initializes a new Index.
func NewIndex() *Index {
	return &Index{
		Index: make(map[string]map[Comic]bool),
	}
}

// Tokenize splits a document into terms.
func Tokenize(content string) []string {
	return strings.Fields(content)
}

// CleanContent removes non-alphanumeric characters from the content.
func CleanContent(content string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	cleanedContent := re.ReplaceAllString(strings.Trim(content, ""), "")
	return cleanedContent
}

// BuildIndex builds the index from the provided comics.
func (idx *Index) BuildIndex(comics map[int]XKCDSearchResult) {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	for _, comic := range comics {
		idx.AddDocument(comic)
	}
}

func (idx *Index) AddDocument(doc XKCDSearchResult) {
	tokens := Tokenize(doc.Transcript)
	for _, token := range tokens {
		normalizedToken := Normalize(token)
		if _, found := idx.Index[normalizedToken]; !found {
			idx.Index[normalizedToken] = map[Comic]bool{}
		}
		idx.Index[normalizedToken][Comic(doc.Num)] = true
	}
}

// Normalize converts a token to lowercase.
func Normalize(token string) string {
	return strings.ToLower(CleanContent(token))
}

func SearchComic(comicNum string) (*XKCDSearchResult, error) {
	url := strings.Join([]string{xkcdURL, comicNum, "info.0.json"}, "/")
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result XKCDSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	return &result, nil
}

func loadComics(filename string) (*ComicIndex, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var comics *ComicIndex
	if err := json.Unmarshal(data, &comics); err != nil {
		return nil, err
	}

	return comics, nil
}

func StoreComics(data ComicIndex) {
	var comicsIndex ComicIndex
	// Check if the file exists and load existing data if it does
	if _, err := os.Stat(filePath); err == nil {
		if comics, err := loadComics(filePath); err == nil {
			comicsIndex = *comics
		} else {
			log.Printf("Failed to load existing comics: %v\n", err)
			comicsIndex = data
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

	jsonData, jsonErr := json.MarshalIndent(comicsIndex, "", " ")
	if jsonErr != nil {
		log.Fatalf("Failed to marshal JSON: %v\n", jsonErr)
	}

	dirName := "comics"
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func Fetch(terms []string) (map[int]XKCDSearchResult, int) {
	total := 0
	comics := map[int]XKCDSearchResult{}
	for _, comic := range terms {
		data, err := SearchComic(comic)
		if err != nil {
			log.Fatalf("Failed to fetch data from comic %s", comic)
		}
		comics[data.Num] = *data
		total++
	}
	return comics, total
}

func Search(term string, comicIndex ComicIndex) {
	normalizedTerm := Normalize(term)
	docPositions, found := comicIndex.Index[normalizedTerm]
	if !found {
		return
	}
	for docID := range docPositions {

		fmt.Printf("URL: https://xkcd.com/%d/\n", comicIndex.Comics[int(docID)].Num)
		fmt.Printf("Transcript: %s\n\n", comicIndex.Comics[int(docID)].Transcript)
		fmt.Println("==========================================")
	}
}

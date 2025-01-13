package main

import (
	"fmt"
	"hash/crc32"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Url struct {
	ShortUrl string // Shortened version of the URL
	LongUrl  string // Original long URL
}
type Store struct {
	Urls map[string]Url
}

func NewStore() *Store {
	return &Store{}
}

func hashUrl(longUrl string) string {
	table := crc32.MakeTable(crc32.IEEE)
	hash := crc32.Checksum([]byte(longUrl), table)
	return fmt.Sprintf("%08x", hash) // Ensure it's zero-padded to 8 characters
}

func (s *Store) AddUrl(url string) string {
	shortUrl := hashUrl(url)
	tempUrl := Url{
		ShortUrl: shortUrl,
		LongUrl:  url,
	}

	if s.Urls == nil {
		s.Urls = make(map[string]Url)
	}

	s.Urls[shortUrl] = tempUrl
	return shortUrl
}

func (s *Store) RemoveUrl(shortUrl string) bool {
	if _, exists := s.Urls[shortUrl]; !exists {
		return false
	}

	delete(s.Urls, shortUrl)
	return true
}

func (s *Store) GetUrl(shortUrl string) string {
	if _, exists := s.Urls[shortUrl]; !exists {
		log.Fatal("shortUrl is not exists")
		return ""
	}
	return s.Urls[shortUrl].LongUrl
}

type URLRequest struct {
	URL string `json:"url"` // Bind the "url" field from the JSON payload
}

func main() {
	r := gin.Default()
	store := NewStore()
	r.GET(":url", func(ctx *gin.Context) {
		shortUrl := ctx.Params.ByName("url")
		url := store.GetUrl(shortUrl)
		if url == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Cannot found this url"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"url": url})
	})

	r.POST("/", func(ctx *gin.Context) {
		var request URLRequest
		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		longUrl := request.URL
		shortUrl := store.AddUrl(longUrl)
		ctx.JSON(http.StatusOK, gin.H{"url": shortUrl})
	})

	r.Run(":8080")
}

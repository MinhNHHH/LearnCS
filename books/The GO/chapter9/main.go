package main

import (
	"io/ioutil"
	"net/http"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func main() {
	// m := memo.New(httpGetBody)
	// for url := range incomingURLs() {
	// 	start := time.Now()
	// 	value, err := m.Get(url)
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	fmt.Printf("%s, %s, %d bytes\n",
	// 		url, time.Since(start), len(value.([]byte)))
	// }
}

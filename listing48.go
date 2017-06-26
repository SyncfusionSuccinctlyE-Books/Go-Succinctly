// Code listing 48: Not available in playground

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type WebPage struct {
	URL  string
	Size int
}

type WebPages []WebPage

// implementing the sort.Interface interface in WebPages
func (slice WebPages) Len() int {
	return len(slice)
}

func (slice WebPages) Less(i, j int) bool {
	// Sort of size of response in descending order
	return slice[i].Size > slice[j].Size
}

func (slice WebPages) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// method for adding a new WebPage element to WebPages
func (wp *WebPages) addElement(page WebPage) {
	*wp = append(*wp, page)
}

// called as a goroutine to retrieve the length of each web page
func getWebPageLength(url string, resultsChannel chan WebPage) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// get the size of the response body
	size, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// populate the WebPage struct and add it to the channel
	var page WebPage
	page.URL = url
	page.Size = len(size)
	resultsChannel <- page
}

func main() {
	urls := []string{
		"http://www.syncfusion.com",
		"http://www.google.com",
		"http://www.microsoft.com",
		"http://www.apple.com",
		"http://www.golang.org",
	}

	// create a channel
	resultsChannel := make(chan WebPage)

	// call a goroutine to read each web page simultaneously
	for _, url := range urls {
		/* initiate a new goroutine for each URL
		so that we can analyze them concurrently */
		go getWebPageLength(url, resultsChannel)
	}

	// store each WebPage result in WebPages
	results := new(WebPages)
	for range urls {
		result := <-resultsChannel
		results.addElement(result)
	}

	// sort using the implementation of sort.Interface in WebPages
	sort.Sort(results)

	// display the results to the user
	for i, page := range *results {
		fmt.Printf("%d. %s: %d bytes.\n", i+1, page.URL,
                                                     page.Size)
	}

}
package scraper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:12.0) Gecko/20100101 Firefox/12.0"

type Scraper interface {
	CreateDocFromURL(url string) error
	FetchDepartingFlights() error
	Info()
}

func CreateDocFromURL(url string) (*goquery.Document, error) {
	client := NewHttpClient(nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// Avoid getting cached pages
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("User-Agent", UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error to execute request: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to construct document from server response: %v", err)
	}

	return doc, nil
}

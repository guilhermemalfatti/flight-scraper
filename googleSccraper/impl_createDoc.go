package googleScraper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gmalfatti/flight-scraper/scraper"
)

func (s *GoogleScraper) CreateDocFromURL(url string) error {
	client := scraper.NewHttpClient(nil)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// Avoid getting cached pages
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("User-Agent", scraper.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error to execute request: %v", err)
	}
	defer resp.Body.Close()

	s.doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("Unable to construct document from server response: %v", err)
	}

	return nil
}

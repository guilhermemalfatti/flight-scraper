package googleScraper

import (
	"fmt"

	"github.com/gmalfatti/flight-scraper/scraper"
)

func (s *GoogleScraper) CreateDocFromURL(url string) error {

	doc, err := scraper.CreateDocFromURL(url)
	if err != nil {
		return err
	}
	s.doc = doc
	return nil
}

func (s *GoogleScraper) FetchDepartingFlights() error {

	return nil
}

func (s *GoogleScraper) Info() {

	fmt.Println("info from google")
}

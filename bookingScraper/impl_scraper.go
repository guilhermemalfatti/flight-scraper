package bookingScraper

import (
	"fmt"

	"github.com/gmalfatti/flight-scraper/scraper"
)

func (s *BookingScraper) CreateDocFromURL(url string) error {

	doc, err := scraper.CreateDocFromURL(url)
	if err != nil {
		return err
	}
	s.doc = doc
	return nil
}

func (s *BookingScraper) FetchDepartingFlights() error {

	return nil
}

func (s *BookingScraper) Info() {

	fmt.Println("info from booking")
}

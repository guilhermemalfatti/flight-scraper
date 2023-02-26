package bookingScraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gmalfatti/flight-scraper/scraper"
	"github.com/sirupsen/logrus"
)

type BookingScraper struct {
	doc *goquery.Document
	log *logrus.Logger
}

func NewScraper() scraper.Scraper {
	return &BookingScraper{
		log: scraper.NewLogger(),
	}
}

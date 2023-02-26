package googleScraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gmalfatti/flight-scraper/scraper"
	"github.com/sirupsen/logrus"
)

type GoogleScraper struct {
	doc *goquery.Document
	log *logrus.Logger
}

func NewScraper() scraper.Scraper {
	return &GoogleScraper{
		log: scraper.NewLogger(),
	}
}

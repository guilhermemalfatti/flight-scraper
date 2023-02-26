package scraper

const UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:12.0) Gecko/20100101 Firefox/12.0"

type Scraper interface {
	CreateDocFromURL(url string) error

	Info()
}

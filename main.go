package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// initialize a data structure to keep the scraped data
type LoginPage struct {
	Login, Bouton string
}

type MDPPage struct {
	MDP, Bouton string
}

type MainAccountPage struct {
	Bouton string
}

func main() {
	fmt.Println("Bank Scrapping")

	// var lp LoginPage

	// instantiate a new collector object
	c := colly.NewCollector(
		colly.AllowedDomains("mon.cmb.fr"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// triggered when the scraper encounters an error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	// fired when the server responds
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// triggered when a CSS selector matches an element
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// printing all URLs associated with the <a> tag on the page
		fmt.Println("%v", e.Attr("href"))
	})

	// triggered once scraping is done (e.g., write the data to a CSV file)
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

	c.OnHTML("#input", func(e *colly.HTMLElement) {
		// ... scraping logic
	})

	c.Visit("https://mon.cmb.fr/auth/login")
}

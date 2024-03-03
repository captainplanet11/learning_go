package main

import (
	"encoding/csv"
	"log"
	"os"
	"sync"

	"github.com/gocolly/colly"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	// initializing the slice of structs to store the data to scrape
	var pokemonProducts []PokemonProduct

	// creating a new Colly instance
	c := colly.NewCollector()

	// synchronization primitive to wait for scraping to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// visiting the target page
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	// wait for scraping to finish before writing to the CSV file
	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	// handle errors while scraping
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		wg.Done()
	})

	// visiting the target page
	c.Visit("https://scrapeme.live/shop/")

	// wait for scraping to finish
	wg.Wait()

	// opening the CSV file
	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}
	if err := writer.Write(headers); err != nil {
		log.Fatalln("Error writing headers to CSV:", err)
	}

	// writing each Pokemon product as a CSV row
	for _, pokemonProduct := range pokemonProducts {
		// converting a PokemonProduct to an array of strings
		record := []string{
			pokemonProduct.url,
			pokemonProduct.image,
			pokemonProduct.name,
			pokemonProduct.price,
		}

		// adding a CSV record to the output file
		if err := writer.Write(record); err != nil {
			log.Println("Error writing record to CSV:", err)
		}
	}

	// flush and check for errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatalln("Error flushing writer:", err)
	}
}

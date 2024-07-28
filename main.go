package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lmcewen9/shopify_crd/scraper"
)

func main() {
	s, err := scraper.fetchShopify(scraper.Configuration{
		URL: os.Getenv("SHOPIFY_URL"),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(s)
}
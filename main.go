package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lmcewen9/shopify_crd/scraper"
)

func main() {
	os.Setenv("SHOPIFY_URL", "https://store.taylorswift.com")
	page := 1
	for {
		s, err := scraper.FetchShopify(&scraper.Configuration{
			URL: os.Getenv("SHOPIFY_URL"),
		}, page)

		if s == "" {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(s)
		page++
	}
}

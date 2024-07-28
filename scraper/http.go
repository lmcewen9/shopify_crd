package scraper

import (
	"encoding/json"
	"log"
	"net/http"
)

type Configuration struct{
	URL		string
}

func fetchShopify(config Configuration) (string, error){
	res, err := http.Get(config.URL + "/products.json")

	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()

	var pResp ProductResponse

	if err := json.NewDecoder(res.Body).Decode(&pResp); err != nil{
		log.Fatal(err)
	}

	return pResp.Products.TextOutput(), nil

}
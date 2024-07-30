package scraper

import (
	"encoding/json"
	"fmt"

	//"io"
	"log"
	"net/http"
)

type Configuration struct {
	URL string
}

func FetchShopify(config *Configuration) (string, error) {
	var url string = fmt.Sprintf("%s/products.json", config.URL)
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var pResp ProductResponse

	if err := json.NewDecoder(res.Body).Decode(&pResp); err != nil {
		log.Fatal(err)
	}

	var ret string
	fmt.Println(len(pResp.Products))

	for i := 0; i < len(pResp.Products); i++ {
		ret += pResp.Products[i].TextOutput(config)
	}

	return ret, nil

}

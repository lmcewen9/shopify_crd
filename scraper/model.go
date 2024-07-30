package scraper

import (
	"fmt"
)

type ProductResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	Category    string     `json:"product_type"`
	Name        string     `json:"title"`
	Variant     []Variants `json:"variants"`
	Product_URL string     `json:"handle"`
}

type Variants struct {
	Price string `json:"price"`
}

func (s Product) TextOutput(config *Configuration) string {
	p := fmt.Sprintf(
		"Name: %s\nCategory: %s\nPrice: %s\nURL: %s/products/%s\n\n",
		s.Name, s.Category, s.Variant[0].Price, config.URL, s.Product_URL)
	return p
}

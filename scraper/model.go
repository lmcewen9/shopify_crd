package scraper

import (
	"fmt"
)

type ProductResponse []struct {
	Products	Product	`json:"products"`
}

type Product []struct {
	Category	string	`json:"product_type"`
	Name		string	`json:"title"`
	Variant		string	`json:"variants"`
	Price		string	`json:"price"`
	Product_URL	string	`json:"handle"`
}

func (s Product) TextOutput() string {
	p := fmt.Sprintf(
		"Category: %s\nName: %s\nVariant: %s\nPrice: %s\n URL: %s/products/%s",
		s[0].Category, s[0].Name, s[0].Variant, s[0].Price, "", s[0].Product_URL)
	return p
}
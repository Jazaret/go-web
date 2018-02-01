package viewmodel

import "github.com/jazaret/go-web/model"

type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply",
		Active:   "Shop",
		Products: []Product{},
	}

	for _, p := range products {
		result.Products = append(result.Products, productToVM(&p))
	}

	return result
}

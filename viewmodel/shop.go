package viewmodel

import (
	"fmt"

	"github.com/jazaret/go-web/model"
)

type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:  "Lemonade Stand Supply - Shop",
		Active: "shop",
	}

	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categorytoVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}

	return result

	juiceCategory := Category{
		URL:      "/shop_details",
		ImageURL: "lemon.png",
		Title:    "Juices and Mixes",
		Description: `Explore our wide assortment of juices and mixes expected by today's lemonade stand clientelle. Now featuring a full line of organic juices that are guaranteed to be obtained from trees that	have never been treated with pesticides or artificial fertilizers.`,
	}

	supplyCategory := Category{
		URL:           ".",
		ImageURL:      "kiwi.png",
		Title:         "Cups, Straws, and Other Supplies",
		Description:   `From paper cups to bio-degradable plastic to straws and napkins, LSS is your source for the sundries that keep your stand running smoothly.`,
		IsOrientRight: true,
	}

	advertizeCategory := Category{
		URL:         ".",
		ImageURL:    "pineapple.png",
		Title:       "Signs and Advertising",
		Description: `Sure, you could just wait for people to find your stand along the side of the road, but if you want to take it to the next level, our premium line of advertising supplies.`,
	}

	result.Categories = []Category{juiceCategory, supplyCategory, advertizeCategory}

	return result
}

func categorytoVM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}

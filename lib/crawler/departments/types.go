package departments

import (
	"amazon/lib/types"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetTypes(subcategory types.SubCategory) types.SubCategory {
	collector := colly.NewCollector()

	collector.OnHTML("#departments", func(element *colly.HTMLElement) {
		element.ForEach(".a-spacing-micro.s-navigation-indent-2", func(_ int, element *colly.HTMLElement) {
			var Type types.Type
			Type.ID = primitive.NewObjectID()
			Type.Url = "https://amazon.com" + element.ChildAttr("a", "href")
			Type.Title = element.ChildText("a")
			subcategory.Types = append(subcategory.Types, Type)
		})
	})
	
	collector.Visit(subcategory.Url)
	return subcategory
}
package departments

import (
	"amazon/lib/types"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetSubTypes(Type types.Type) types.Type {
	collector := colly.NewCollector()

	collector.OnHTML("#departments", func(element *colly.HTMLElement) {
		element.ForEach(".a-spacing-micro.s-navigation-indent-2", func(_ int, element *colly.HTMLElement) {
			var subType types.SubType
			subType.ID = primitive.NewObjectID()
			subType.Url = "https://amazon.com" + element.ChildAttr("a", "href")
			subType.Title = element.ChildText("a")
			Type.SubTypes = append(Type.SubTypes, subType)
		})
	})
	
	collector.Visit(Type.Url)
	return Type
}
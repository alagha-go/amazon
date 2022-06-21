package departments

import (

	"github.com/gocolly/colly"
)


func (subcategory *SubCategory) SetTypes(){
	collector := colly.NewCollector()

	collector.OnHTML("#departments", func(element *colly.HTMLElement) {
		element.ForEach(".a-spacing-micro.s-navigation-indent-2", func(_ int, element *colly.HTMLElement) {
			var Type Type
			Type.Url = "https://amazon.com" + element.ChildAttr("a", "href")
			Type.Title = element.ChildText("a")
			subcategory.Types = append(subcategory.Types, Type)
		})
	})

	if len(subcategory.Types) == 0 {
		collector.OnHTML("#s-refinements", func(element *colly.HTMLElement) {
			element.ForEach(".a-spacing-micro.apb-browse-refinements-indent-2", func(_ int, element *colly.HTMLElement) {
				var Type Type
				Type.Url = "https://amazon.com" + element.ChildAttr("a", "href")
				Type.Title = element.ChildText("a")
				subcategory.Types = append(subcategory.Types, Type)
			})
		})
	}
	
	collector.Visit(subcategory.Url)
}
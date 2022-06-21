package departments

import (

	"github.com/gocolly/colly"
)



func (department *Department) SetCategories() {
	for index := range department.Categories {
		department.Categories[index].SetSubCategories()
	}
}

func (category *Category) SetSubCategories() {
	collector := colly.NewCollector()

	
	collector.OnHTML("#departments", func(element *colly.HTMLElement) {
		element.ForEach(".a-spacing-micro.s-navigation-indent-2", func(_ int, element *colly.HTMLElement) {
			var subCategory SubCategory
			subCategory.Url = "https://amazon.com" + element.ChildAttr("a", "href")
			subCategory.Title = element.ChildText("a")
			category.SubCategories = append(category.SubCategories, subCategory)
		})
	})

	if len(category.SubCategories) == 0 {
		collector.OnHTML("#s-refinements", func(element *colly.HTMLElement) {
			element.ForEach(".a-spacing-micro.apb-browse-refinements-indent-2", func(_ int, element *colly.HTMLElement) {
				var subCategory SubCategory
				subCategory.Url = "https://amazon.com" + element.ChildAttr("a", "href")
				subCategory.Title = element.ChildText("a")
				category.SubCategories = append(category.SubCategories, subCategory)
			})
		})
	}
	
	collector.Visit(category.Url)
}
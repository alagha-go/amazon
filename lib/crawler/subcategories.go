package crawler

import (
	"amazon/lib/types"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetSubCategories(departments []types.Department) []types.Department {
	var NewDepartments []types.Department
	for _, department := range departments {
		for index, category := range department.Categories {
			department.Categories[index].SubCategories = CollectSubCategories(category)
		}
		NewDepartments = append(NewDepartments, department)
	}

	return NewDepartments
}

func CollectSubCategories(category types.Category) []types.SubCategory {
	collector := colly.NewCollector()
	var SubCategories []types.SubCategory

	collector.OnHTML("#departments", func(element *colly.HTMLElement) {
		element.ForEach(".a-spacing-micro.s-navigation-indent-2", func(_ int, element *colly.HTMLElement) {
			var subCategory types.SubCategory
			subCategory.ID = primitive.NewObjectID()
			subCategory.Url = element.ChildAttr("a", "href")
			subCategory.Title = element.ChildText("a")
			SubCategories = append(SubCategories, subCategory)
		})
	})

	return SubCategories
}
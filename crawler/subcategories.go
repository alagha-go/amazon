package crawler

import (
	// "time"

	"fmt"
	"strings"

	"github.com/gocolly/colly"
)



func (department *Department) SetCategories() {
	for index := range department.Categories {
		department.Categories[index].SetSubCategories()
	}
}

func (category *Category) SetSubCategories() {
	collector := colly.NewCollector()

	err := collector.SetProxy(ProxyUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})

	collector.OnHTML("body", func(element *colly.HTMLElement) {
		element.ForEach("li", func(_ int, element *colly.HTMLElement) {
			if element.Attr("class") == "a-spacing-micro s-navigation-indent-2" {
				element.ForEachWithBreak("a", func(_ int, element *colly.HTMLElement) bool {
					off := element.Attr("data-routing")
					if off == "off" {
						var subCategory SubCategory
						subCategory.Url = "https://amazon.com" + element.Attr("href")
						subCategory.Title = element.Text
						subCategory.Title = strings.ReplaceAll(subCategory.Title, "\n", "")
						subCategory.Title = strings.ReplaceAll(subCategory.Title, "  ", " ")
						subCategory.Title = strings.TrimSpace(subCategory.Title)
						category.SubCategories = append(category.SubCategories, subCategory)
						return false
					}
					return true
				})
			}
		})
	})
	
	collector.Visit(category.Url)
	Reload()
}